package physics

type MagnetConfig struct {
	Radius         float64
	Pull           float64
	FollowStrength float64
}

type PhysicsConfig struct {
	GravityFalloff float64
	ShipFriction   float64
	MaxShipSpeed   float64
}
