package render

/***********************
 * Material
 ************************/
// Material defines how a material scatter light
type Material interface {
	Scatter(r *Ray3, rec *HitRecord) (wasScattered bool, attenuation *Color, scattered *Ray3)
}
