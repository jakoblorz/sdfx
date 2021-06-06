package render

// Camera defines the api to use the camera => computes a ray
// Implementation note: the camera is provided its own source of random number for 2 reasons
// 	1. allow to abstract random to make it non random if necessary for testing
//  2. using (global) rand.Float() turns out to be a major slowdown when using multiple goroutines as
// 		due to obvious reasons, it needs to be synchronized => can use a non synchronized version
type Camera interface {
	Ray(rnd Rnd, u, v float64) *Ray3
}
