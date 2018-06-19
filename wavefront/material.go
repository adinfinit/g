package wavefront

import "github.com/adinfinit/g"

// https://en.wikipedia.org/wiki/Wavefront_.obj_file
// http://paulbourke.net/dataformats/mtl/
type Material struct {
	Name string

	Ambient struct {
		Color   g.Vec3 // Ka
		Texture string // map_Ka
	}

	Diffuse struct {
		Color   g.Vec3 // Kd
		Texture string // map_Kd
	}
	Specular struct {
		Color    g.Vec3  // Ks
		Exponent float32 // Ns
		Texture  string  // map_Ks
	}

	Emissive struct {
		Color g.Vec3 // Ke
	}

	Illumination struct {
	}

	Alpha struct {
		Value   float32 // d=0 or Tr=1 is opaque  (Tr = 1 - d)
		Texture string  // map_d
	}

	Displacement struct {
		Texture string // disp
	}
	Bump struct {
		Texture string // map_bump or bump
	}
	Decal struct {
		Texture string // decal
	}
}
