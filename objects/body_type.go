package objects

type BodyType string

const (
	BodyPlanet      BodyType = "planet"
	BodyDestination BodyType = "destination"
	BodyStar        BodyType = "star"
	BodyAsteroid    BodyType = "asteroid"
	BodyComet       BodyType = "comet"
	BodyBlackHole   BodyType = "blackhole"
	BodyWormHole    BodyType = "wormhole"
)
