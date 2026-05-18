package game

import (
	"math"
	"math/rand/v2"

	"github.com/FBascou/rocket_game/entities"
	"github.com/FBascou/rocket_game/shared"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

// Initializes ebiten.Image images for ship, planets, minerals, etc.
func (game *Game) InitializeAssets() {
	game.Assets.ShipImage = loadImage("assets/ship.png")
}

// This should be called when:
// resetting ship only,
// restore minerals to original state,
// not regenerating planets,
// restart the same level if player crashed less than 5 times
func (game *Game) resetLevel() {
	game.GameState = shared.StateAiming
	game.Ship = game.generateShip()
	game.Planets = deepClonePlanets(game.InitialPlanets)
	game.CollectedMinerals = 0
}

func (game *Game) resetShipAfterCrash() {
	game.GameState = shared.StateAiming
	game.Ship = game.generateShip()
}

func (game *Game) generateShip() entities.Ship {
	return entities.Ship{
		X:               float64(game.ScreenWidth) / 2,
		Y:               float64(game.ScreenHeight) - 50,
		Rotation:        0,
		MagnetRadius:    game.GameConfig.MagnetRadius,
		CollisionRadius: 6,
	}
}

func (game *Game) generatePlanets() []entities.Planet {
	// min 1 planet, max 3 planets (not including destination panet)
	numberOfPlanets := 1 + rand.IntN(3)

	screenW := float64(game.ScreenWidth)
	screenH := float64(game.ScreenHeight)

	topMargin := 80.0
	bottomMargin := 120.0

	usableHeight := screenH - topMargin - bottomMargin
	// value that affexts the vertical distance between planets
	// add a verticalSpacingMultiplier := 1.4
	// stepY := (usableHeight / float64(numberOfPlanets + 1)) * verticalSpacingMultiplier
	stepY := usableHeight / float64(numberOfPlanets+1)

	centerX := screenW / 2
	// static value that affects the horizontal distance between planets
	xVariation := 60.0

	planets := []entities.Planet{}

	for i := 0; i < numberOfPlanets; i++ {
		// min 2 minerals, max 4 minerals
		numberOfMinerals := 2 + rand.IntN(3)

		var planet entities.Planet

		// boolean that guarantees no overlap between planets
		validPlanet := false

		// if true, create a planet on the calculated position
		for !validPlanet {
			// a bit of planet x randomness with a left-path (more zig-zag)
			// direction alternates left/right (offset)
			direction := 1.0
			if i%2 == 0 {
				direction = -1.0
			}
			// planet x position = center + left/right offset
			// 20 + rand.Float64() (0 to 1) * 40 = random number between 20 and 60 px
			// currently creates a 20-60px left/right offset
			x := centerX + direction*(20+rand.Float64()*40)

			// a bit of planet x randomness (more linear than above)
			// x := centerX + (rand.Float64() * 2 - 1) * xVariation

			// controlled planet y variation
			// y := bottomMargin + float64(i) * stepY
			y := bottomMargin + float64(i+1)*stepY

			planet = entities.Planet{
				X: x,
				Y: y,
				// physical size
				Size: 30,
				// gravity field size
				GravityRadius: 80,
				// pull force
				GravityStrength: 0.5,
				BodyType:        entities.BodyDestination,
			}

			validPlanet = true

			// collision check: prevent overlap, planets don't touch each other
			for _, existing := range planets {
				// horizontal difference between 2 planets
				dx := existing.X - planet.X
				// vertical difference between 2 planets
				dy := existing.Y - planet.Y
				// distance between 2 planets (vector between two objects)
				dist := math.Sqrt(dx*dx + dy*dy)

				// adding planet padding between each other
				// are the circles/planets touching (+ padding)
				// this controls how close planets are allowed to spawn
				if dist < existing.Size+planet.Size+30 {
					validPlanet = false
					break
				}
			}
		}

		planet.Minerals = game.generateMinerals(planet, numberOfMinerals)
		planets = append(planets, planet)
	}

	var destinationPlanet entities.Planet
	validDestination := false

	// if there's not enough space in the window, there could be an infinite loop
	// this is a safety counter to avoid an infinite loop if
	attempts := 0

	for !validDestination && attempts < 50 {
		attempts++
		// centerX + ((rand between -1 and +1) * 60) = centerX +- 60
		// centers destination planet with horizontal skew
		x := centerX + (rand.Float64()*2-1)*xVariation

		// gives space between planets and top edge
		y := topMargin + 20

		destinationPlanet = entities.Planet{
			X:               x,
			Y:               y,
			Size:            50,
			GravityRadius:   140,
			GravityStrength: 2.5,
			IsDestination:   true,
			BodyType:        entities.BodyDestination,
		}

		validDestination = true

		for _, existing := range planets {
			dx := existing.X - destinationPlanet.X
			dy := existing.Y - destinationPlanet.Y
			dist := math.Sqrt(dx*dx + dy*dy)

			if dist < existing.Size+destinationPlanet.Size+40 {
				validDestination = false
				break
			}
		}
	}

	planets = append(planets, destinationPlanet)

	return planets
}

func (game *Game) generateMinerals(planet entities.Planet, count int) []entities.Mineral {
	minerals := []entities.Mineral{}
	minDistance := 15.0

	for len(minerals) < count {
		angle := rand.Float64() * 2 * math.Pi
		// uniform distribution of minerals across the circle
		// minerals distribute evenly over the whole planet area
		radius := math.Sqrt(rand.Float64()) * planet.Size

		x := planet.X + radius*math.Cos(angle)
		y := planet.Y + radius*math.Sin(angle)

		valid := true

		for _, mineral := range minerals {
			dx := mineral.X - x
			dy := mineral.Y - y
			dist := math.Sqrt(dx*dx + dy*dy)

			if dist < minDistance {
				valid = false
				break
			}
		}

		if valid {
			minerals = append(minerals, entities.Mineral{X: x, Y: y})
		}
	}

	return minerals
}

func (game *Game) collectMinerals() {
	for pi := range game.Planets {
		planet := &game.Planets[pi]

		for mi := range planet.Minerals {
			mineral := &planet.Minerals[mi]

			if mineral.Collected {
				continue
			}

			// apply velocity and move mineral
			mineral.X += mineral.VX
			mineral.Y += mineral.VY

			// damping (prevents infinite speed)
			mineral.VX *= game.GameConfig.MineralDamping
			mineral.VY *= game.GameConfig.MineralDamping

			// recalculate distance (NEW position)
			dx := game.Ship.X - mineral.X
			dy := game.Ship.Y - mineral.Y

			distance := math.Sqrt(dx*dx + dy*dy)

			// ship's magnet pull strength and mineral's speed
			if distance > 0 && distance < game.Ship.MagnetRadius {
				nx := dx / distance
				ny := dy / distance

				shipSpeed := math.Sqrt(game.Ship.VX*game.Ship.VX + game.Ship.VY*game.Ship.VY)

				// base magnet pull with smooth acceleration, becomes stronger when closer (far=weak pull, near=strong pull)
				strength := 1 - (distance / game.Ship.MagnetRadius) // 0 to 1

				// ship's magnetic pull (numbers are placeholders)
				pull := (game.GameConfig.MagnetPull + shipSpeed*0.1) * strength

				mineral.VX += nx * pull
				mineral.VY += ny * pull

				// to prevent mineral lag when getting pulled by a fast travelling ship, faster ship = stronger magnet pul
				// 0.2 is the follow through strength to the ship, 0 = minerals will lag behind a faster ship, X = minerals will reach ship fast
				follow := game.GameConfig.MagnetFollowStrength
				mineral.X += dx * follow
				mineral.Y += dy * follow

				// dynamic max speed (3 is a placeholder for now)
				mineralMaxSpeed := shipSpeed + 3
				mineralSpeed := math.Sqrt(mineral.VX*mineral.VX + mineral.VY*mineral.VY)

				if mineralSpeed > mineralMaxSpeed {
					scale := mineralMaxSpeed / mineralSpeed
					mineral.VX *= scale
					mineral.VY *= scale
				}
			}

			// collect mineral
			if distance < 10 {
				mineral.Collected = true
				game.CollectedMinerals++
			}
		}
	}
}

func (game *Game) checkWin() {
	destinationPlanet := game.getDestinationPlanet()

	dx := destinationPlanet.X - game.Ship.X
	dy := destinationPlanet.Y - game.Ship.Y

	distance := math.Sqrt(dx*dx + dy*dy)

	if distance < destinationPlanet.Size {
		game.getWinResult()
	}
}

func (game *Game) collectStars() {
	totalMinerals := getTotalMineralsInLevel(game.Planets)
	allMineralsCollected := game.CollectedMinerals == totalMinerals
	is3stars := allMineralsCollected
	is2stars := game.CrashCount == 0
	is1star := game.CollectedMinerals == 0 && game.CrashCount != 0

	if is3stars {
		game.Stars = 3 // All minerals + all lives
	} else if is2stars {
		game.Stars = 2 // Not all minerals + all lives
	} else if is1star {
		game.Stars = 1 // No minerals + not all lives
	} else {
		game.Stars = 0 // No minerals + 0 lives
	}
}

func (game *Game) updateMenu() {
	if !inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		return
	}

	mx, my := ebiten.CursorPosition()

	// replay button
	if isPointInsideRect(mx, my, 100, 380, 200, 40) {
		game.resetLevel()
		game.LevelMenuState = shared.LevelMenuClosed
	}

	// next level button or continue button (keep playing)
	if isPointInsideRect(mx, my, 100, 440, 200, 40) {
		// go next level
		if game.LevelMenuState == shared.LevelMenuWin {
			game.nextLevel()
			game.LevelMenuState = shared.LevelMenuClosed
			return
		}

		// keep playing after crash/pause
		if game.LevelMenuState == shared.LevelMenuCrash ||
			game.LevelMenuState == shared.LevelMenuPause {
			game.LevelMenuState = shared.LevelMenuClosed
			return
		}
	}
}
