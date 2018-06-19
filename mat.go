package g

type Mat4 struct {
	M00, M01, M02, M03 float32
	M10, M11, M12, M13 float32
	M20, M21, M22, M23 float32
	M30, M31, M32, M33 float32
}

const Mat4Size = 4 * 4 * 4

func (mat *Mat4) Ptr() *float32 { return &mat.M00 }

func Perspective(fovy, aspect, near, far float32) Mat4 {
	f := 1.0 / Tan(fovy/2.0)
	return Mat4{
		f / aspect, 0, 0, 0,
		0, f, 0, 0,
		0, 0, -(near + far) / (far - near), -1,
		0, 0, -2.0 * far * near / (far - near), 0,
	}
}

func LookAtV(eye, center, up Vec3) Mat4 {
	f := center.Sub(eye).Normalize().Neg()
	s := f.Cross(up.Normalize()).Normalize()
	u := s.Cross(f)
	return Mat4{
		s.X, u.X, f.X, 0,
		s.Y, u.Y, f.Y, 0,
		s.Z, u.Z, f.Z, 0,
		-eye.X, -eye.Y, -eye.Z, 1,
	}
}
