package main

import (
	"image/color"
	"math"
	"math/rand/v2"

	"github.com/hajimehoshi/ebiten/v2"
)

type Assets struct {
	ShipImage    *ebiten.Image
	PlanetImage  *ebiten.Image
	MineralImage *ebiten.Image
}

type Level struct {
	screenWidth  int
	screenHeight int
} 

type Game struct {
	Assets 						Assets
	Level 						Level
	Ship    					Ship
	InitialPlanets 		[]Planet
	Planets 					[]Planet
	CollectedMinerals int
	CrashCount				int
	Lives 						int
	Stars 						int
	Launched 					bool
	Crashed  					bool
	Won 	   					bool
}

// Initializes ebiten.Image images for ship, planets, minerals, etc.
func (game *Game) initializeAssets() {
	game.Assets.ShipImage = loadImage("assets/ship.png")
}

func (game *Game) initializeAssetOutlines() {
	game.Ship.MagnetRadiusOutline = createCircleImage(int(game.Ship.MagnetRadius), color.RGBA{0, 255, 255, 255})

	for index, planet := range game.Planets {
    if planet.IsDestination {
			game.Planets[index].GravityOutline =
				createCircleImage(int(game.Planets[index].Gravity), color.RGBA{0, 0, 255, 255})
    } else {
			game.Planets[index].GravityOutline =
				createCircleImage(int(game.Planets[index].Gravity), color.RGBA{255, 0, 0, 255})
    }
	}
}

// This should be called when: 
// starting the next/new level,
// restarting entire game
func (game *Game) initializeNewLevel() {
	game.Ship = game.generateShip()
	game.Planets = game.generatePlanets()
	game.Launched = false
	game.Won = false
	game.initializeAssetOutlines() 
}

// This should be called when:
// resetting ship only, 
// restore minerals to original state, 
// not regenerating planets,
// restart the same level if player crashed less than 5 times
func (game *Game) restartLevel() {
		game.Ship = game.generateShip()
		game.Launched = false
		game.Won = false
		game.initializeAssetOutlines()
}

func (game *Game) generateShip() Ship {
	return Ship{
		X: float64(dimensionWidth) / 2,
		Y: float64(dimensionHeight) - 50,
		Rotation: 0,
		MagnetRadius: 60,
	}
}

func (game *Game) generatePlanets() []Planet {
	// min 1 planet, max 3 planets (not including destination panet)
	numberOfPlanets := 1 + rand.IntN(3)
	
	screenW := float64(dimensionWidth)
	screenH := float64(dimensionHeight)

	topMargin := 80.0
	bottomMargin := 120.0

	usableHeight := screenH - topMargin - bottomMargin
	// stepY := usableHeight / float64(numberOfPlanets)
	stepY := usableHeight / float64(numberOfPlanets + 1)

	centerX := screenW / 2
	xVariation := 60.0

	planets := []Planet{}

	for i := 0; i < numberOfPlanets; i++ {
		// min 2 minerals, max 4 minerals
		numberOfMinerals := 2 + rand.IntN(3)

		var planet Planet

		// boolean that guarantees no overlap between planets
		validPlanet := false

		// if true, create a planet on the calculated position
		for !validPlanet {
			// a bit of planet x randomness with a left-path (more zig-zag)
			// direction alternates left/right (offset)
			direction := 1.0
			if i % 2 == 0 {
				direction = -1.0
			}
			// planet x position = center + left/right offset
			// 20 + rand.Float64() (0 to 1) * 40 = random number between 20 and 60 px
			x := centerX + direction * (20 + rand.Float64() * 40)

			// a bit of planet x randomness (more linear than above)
			// x := centerX + (rand.Float64() * 2 - 1) * xVariation

			// controlled planet y variation
			// y := bottomMargin + float64(i) * stepY
			y := bottomMargin + float64(i + 1) * stepY

			planet = Planet{
				X: x,
				Y: y,
				Radius: 30,
				Gravity: 80,
			}

			validPlanet = true

			// collision check: prevent overlap, planets don't touch each other
			for _, existing := range planets {
				// horizontal difference between 2 planets
				dx := existing.X - planet.X
				// vertical difference between 2 planets
				dy := existing.Y - planet.Y
				// distance between 2 planets (vector between two objects)
				dist := math.Sqrt(dx * dx + dy * dy)

				// adding planet padding between each other
				// are the circles/planets touching (+ padding)
				if dist < existing.Radius + planet.Radius + 30 {
					validPlanet = false
					break
				}
			}
		}

		planet.Minerals = game.generateMinerals(planet, numberOfMinerals)
		planets = append(planets, planet)
	}
	
	var destinationPlanet Planet
	validDestination := false

	// if there's not enough space in the window, there could be an infinite loop
	// this is a safety counter to avoid an infinite loop if 
	attempts := 0
	
	for !validDestination && attempts < 50 {
		attempts++
		// centerX + ((rand between -1 and +1) * 60) = centerX +- 60
		// centers destination planet with horizontal skew
		x := centerX + (rand.Float64() * 2 - 1) * xVariation

		// gives space between planets and top edge
		y := topMargin + 20

		destinationPlanet = Planet{
			X: x,
			Y: y,
			Radius: 50,
			Gravity: 200, 
			IsDestination: true,
		}

		validDestination = true

		for _, existing := range planets {
			dx := existing.X - destinationPlanet.X
			dy := existing.Y - destinationPlanet.Y
			dist := math.Sqrt(dx * dx + dy * dy)

			if dist < existing.Radius + destinationPlanet.Radius + 40 {
				validDestination = false
				break
			}
		}
	}

	planets = append(planets, destinationPlanet)

	return planets
}

func (game *Game) generateMinerals(planet Planet, count int) []Mineral {
	minerals := []Mineral{}
	minDistance := 15.0

	for len(minerals) < count {
		angle := rand.Float64() * 2 * math.Pi
		radius := rand.Float64() * planet.Radius

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
			minerals = append(minerals, Mineral{X: x, Y: y})
		}
	}

	return minerals
}

// Draw the ship image and its magnet outline
// GeoM order matters: Translate(center of body) > Scale > Rotate > Translate(world position)
func (game *Game) drawShip(screen *ebiten.Image) {
	ship := game.Ship

	w := float64(game.Assets.ShipImage.Bounds().Dx())
	h := float64(game.Assets.ShipImage.Bounds().Dy())
	scale := 0.05

	options := &ebiten.DrawImageOptions{}

	// by default in ebiten, images draw from top-left corner
	// this moves image origin to the center of the image
	// it rotates around the center of the image (image origin) and not top-left corner
	// so it rotates in place instead of swinging in giant circles 
	options.GeoM.Translate(-w / 2, -h / 2)

	// temporary fix for ship size issue
	options.GeoM.Scale(scale, scale)

	// rotate ship
	options.GeoM.Rotate(ship.Rotation)

	// move ship into world position
	options.GeoM.Translate(ship.X, ship.Y)
	

	// draw Ship
	screen.DrawImage(game.Assets.ShipImage, options)

	// draw ship gravity radius outline
	createOutlineImage(screen, ship, ship.MagnetRadius, ship.MagnetRadiusOutline, outlineImageAlpha)
}

// Draw the planet image and its gravity outline
func drawPlanet(screen *ebiten.Image, planet Planet) {
	// draw Planet
		for i := -int(planet.Radius); i < int(planet.Radius); i++ {
			for j := -int(planet.Radius); j < int(planet.Radius); j++ {
				// equation of a circle (x2+y2≤r2)
				if i * i + j * j <= int(planet.Radius * planet.Radius) {

					if planet.IsDestination {
						screen.Set(int(planet.X) + i, int(planet.Y) + j, color.RGBA{0, 0, 255, 255})
					} else {
						screen.Set(int(planet.X) + i, int(planet.Y) + j, color.RGBA{255, 0, 0, 255})
					}

				}
			}
		}
		
		// draw planet gravity radius outline
		createOutlineImage(screen, planet, planet.Gravity, planet.GravityOutline, outlineImageAlpha)
}

func (game *Game) drawPlanets(screen *ebiten.Image) {
	for _, planet := range game.Planets {
		// isDestinationPlanet := index == game.getDestinationPlanetIndex()
		drawPlanet(screen, planet)
	}
}

// Draw the minerals image
func (game *Game) drawMinerals(screen *ebiten.Image) {

	for _, p := range game.Planets {
		for _, m := range p.Minerals {
			if m.Collected {
				continue
			}

			for dx := -2; dx <= 2; dx++ {
				for dy := -2; dy <= 2; dy++ {
					speed := math.Sqrt(m.VX * m.VX + m.VY * m.VY)
					brightness := uint8(math.Min(255, speed * 50))
					col := color.RGBA{brightness, 255, brightness, 255}
					screen.Set(int(m.X) + dx, int(m.Y) + dy, col)
				}
			}
		}
	}
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
			mineral.VX *= 0.95
			mineral.VY *= 0.95
			
			// recalculate distance (NEW position)
			dx := game.Ship.X - mineral.X
			dy := game.Ship.Y - mineral.Y

			distance := math.Sqrt(dx * dx + dy * dy)

			// ship's magnet pull strength and mineral's speed
			if distance > 0 && distance < game.Ship.MagnetRadius {
				nx := dx / distance
				ny := dy / distance

				shipSpeed := math.Sqrt(game.Ship.VX * game.Ship.VX + game.Ship.VY * game.Ship.VY)

				// base magnet pull with smooth acceleration, becomes stronger when closer (far=weak pull, near=strong pull)
				strength := 1 - (distance / game.Ship.MagnetRadius) // 0 to 1

				// ship's magnetic pull (numbers are placeholders)
				pull := (0.3 + shipSpeed * 0.1) * strength		

				mineral.VX += nx * pull
				mineral.VY += ny * pull

				// to prevent mineral lag when getting pulled by a fast travelling ship, faster ship = stronger magnet pul
				// 0.2 is the follow through strength to the ship, 0 = minerals will lag behind a faster ship, X = minerals will reach ship fast
				follow := 0.2
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

	if distance < destinationPlanet.Radius {
			game.Won = true
	}
}

func (game *Game) collectStars() {
	totalMinerals := getTotalMineralsInLevel(game.Planets)
	allMineralsCollected := game.CollectedMinerals == totalMinerals
	is3stars := allMineralsCollected && game.CrashCount == 0
	is2stars := allMineralsCollected || game.CrashCount == 0
	is1star := !allMineralsCollected && game.CrashCount != 0

	if is3stars {
		// All minerals + all lives
		game.Stars = 3 
	} else if is2stars {
		// All minerals + between 0 and 5 lives OR not all minerals + 5 lives 
		game.Stars = 2
	} else if is1star {
		// Not all minerals + between 0 and 5 lives 
		game.Stars = 1
	} else {
		// Not all minerals + 0 lives 
		game.Stars = 0
	}
}