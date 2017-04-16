package wavefront

import (
	"bufio"
	"errors"
	"io"
	"strconv"
	"strings"

	"adinfinitum.ee/g"
)

func parse32(s string) (float32, error) {
	x, err := strconv.ParseFloat(s, 32)
	return float32(x), err
}

func ParseOBJ(rd io.Reader) ([]*Object, []string, error) {
	sc := bufio.NewScanner(rd)
	var objects []*Object
	var materials []string

	var object *Object
	var group *Group

	ensureGroup := func() {
		if object == nil {
			object = &Object{}
			objects = append(objects, object)
		}
		if group == nil {
			object.Groups = append(object.Groups, Group{})
			group = &object.Groups[len(object.Groups)-1]
		}
	}
	nextGroup := func() {
		ensureGroup()
		object.Groups = append(object.Groups, Group{})
		group = &object.Groups[len(object.Groups)-1]
	}

	var err error

	for sc.Scan() {
		line := strings.TrimSpace(sc.Text())
		if line == "" || line[0] == '#' {
			continue
		}

		tokens := strings.Fields(line)
		rest := strings.Join(tokens[1:], " ")

		switch tokens[0] {
		case "mtllib":
			materials = append(materials, rest)
		case "o":
			group = nil
			object = &Object{}
			// TODO: handle multiple names
			object.Name = tokens[1]
			objects = append(objects, object)
		case "g":
			nextGroup()
			// TODO: handle multiple names
			if len(tokens) > 1 {
				group.Name = tokens[1]
			}
		case "v":
			ensureGroup()
			var v g.Vec3
			v.X, err = parse32(tokens[1])
			v.Y, err = parse32(tokens[2])
			v.Z, err = parse32(tokens[3])
			group.Positions = append(group.Positions, v)
		case "vt":
			ensureGroup()
			var v g.Vec3
			v.X, err = parse32(tokens[1])
			v.Y, err = parse32(tokens[2])
			if len(tokens) > 3 {
				v.Z, err = parse32(tokens[3])
			}
			group.TexCoords = append(group.TexCoords, v)
		case "vn":
			ensureGroup()
			var v g.Vec3
			v.X, err = parse32(tokens[1])
			v.Y, err = parse32(tokens[2])
			v.Z, err = parse32(tokens[3])
			group.Normals = append(group.Normals, v)
		case "usemtl":
			ensureGroup()
			group.Material = tokens[1]
		case "s":
			ensureGroup()
			group.Smooth = tokens[1] == "on"
		case "f":
			ensureGroup()
			var face Face
			for _, token := range tokens[1:] {
				var vertex Vertex

				indices := strings.Split(token, "/")
				if len(indices) < 2 {
					indices = append(indices, "", "", "")
				}

				if indices[0] != "" {
					vertex.Position, err = strconv.Atoi(indices[0])
				}
				if indices[1] != "" {
					vertex.TexCoord, err = strconv.Atoi(indices[1])
				}
				if indices[2] != "" {
					vertex.Normal, err = strconv.Atoi(indices[2])
				}

				face.Vertices = append(face.Vertices, vertex)
			}
			group.Faces = append(group.Faces, face)
		default:
			err = errors.New("unhandled value " + tokens[0])
		}
	}

	return objects, materials, err
}
