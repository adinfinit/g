package wavefront

import "adinfinitum.ee/g"

// https://en.wikipedia.org/wiki/Wavefront_.obj_file
// http://paulbourke.net/dataformats/obj/
type Model struct {
	Positions []g.Vec3
	TexCoords []g.Vec3
	Normals   []g.Vec3

	Objects []Object

	MaterialLibs []string
}

type Object struct {
	Name   string
	Groups []Group
}

type Group struct {
	Name     string
	Smooth   bool
	Material string
	Faces    []Face
}

type Face struct {
	Vertices []Vertex
}

type Vertex struct {
	Position int
	TexCoord int
	Normal   int
}
