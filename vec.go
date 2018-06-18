// generated by vec.gen
// DO NOT EDIT

package g

// Vec2 is a 2 element vector
type Vec2 struct{ X, Y float32 }

var (
	Z2 = Vec2{}
	I2 = Vec2{1, 1}
)

func V2(X, Y float32) Vec2 {
	return Vec2{X: X, Y: Y}
}

// Add does a component wise a + b
func (a Vec2) Add(b Vec2) Vec2 {
	return Vec2{X: a.X + b.X, Y: a.Y + b.Y}
}

// Sub does a component wise a - b
func (a Vec2) Sub(b Vec2) Vec2 {
	return Vec2{X: a.X - b.X, Y: a.Y - b.Y}
}

// Mul does a component wise a * b
func (a Vec2) Mul(b Vec2) Vec2 {
	return Vec2{X: a.X * b.X, Y: a.Y * b.Y}
}

// Div does a component wise a / b
func (a Vec2) Div(b Vec2) Vec2 {
	return Vec2{X: a.X / b.X, Y: a.Y / b.Y}
}

// Scale scales each component by s
func (a Vec2) Scale(s float32) Vec2 {
	return Vec2{X: a.X * s, Y: a.Y * s}
}

// Dot returns the dot product of a and b
func (a Vec2) Dot(b Vec2) float32 {
	return (a.X*b.X + a.Y*b.Y)
}

// Len returns vector length
func (a Vec2) Len() float32 {
	return Sqrt(a.X*a.X + a.Y*a.Y)
}

// Len2 returns vector length squared
func (a Vec2) Len2() float32 {
	return (a.X*a.X + a.Y*a.Y)
}

// Normalize returns normalized vector
func (a Vec2) Normalize() Vec2 {
	s := 1.0 / a.Len()
	return Vec2{X: a.X * s, Y: a.Y * s}
}

// Min does an element wise min
func (a Vec2) Min(b Vec2) Vec2 {
	return Vec2{X: Min(a.X, b.X), Y: Min(a.Y, b.Y)}
}

// Max does an element wise max
func (a Vec2) Max(b Vec2) Vec2 {
	return Vec2{X: Max(a.X, b.X), Y: Max(a.Y, b.Y)}
}

// Eq returns whether a == b
func (a Vec2) Eq(b Vec2) bool {
	return (a.X == b.X && a.Y == b.Y)
}

// Lt returns whether a < b
func (a Vec2) Lt(b Vec2) bool {
	return (a.X < b.X && a.Y < b.Y)
}

// Le returns whether a <= b
func (a Vec2) Le(b Vec2) bool {
	return (a.X <= b.X && a.Y <= b.Y)
}

// Gt returns whether a > b
func (a Vec2) Gt(b Vec2) bool {
	return (a.X > b.X && a.Y > b.Y)
}

// Ge returns whether a >= b
func (a Vec2) Ge(b Vec2) bool {
	return (a.X >= b.X && a.Y >= b.Y)
}

// EqAlmost returns whether Abs(a - b) < eps
func (a Vec2) EqAlmost(b Vec2, eps float32) bool {
	return (Abs(a.X-b.X) < eps && Abs(a.Y-b.Y) < eps)
}

// Lerp returns linearly interpolated vector
func (a Vec2) Lerp(b Vec2, p float32) Vec2 {
	return Vec2{X: Lerp(a.X, b.X, p), Y: Lerp(a.Y, b.Y, p)}
}

// Lerp returns linearly interpolated vector
func (a Vec2) LerpClamp(b Vec2, p float32) Vec2 {
	if p < 0 {
		return a
	} else if p > 1 {
		return b
	}
	return Vec2{X: Lerp(a.X, b.X, p), Y: Lerp(a.Y, b.Y, p)}
}

// Clamp returns vector inside min, max
func (a Vec2) Clamp(min, max Vec2) Vec2 {
	return Vec2{X: Clamp(a.X, min.X, max.X), Y: Clamp(a.Y, min.Y, max.Y)}
}

// Clamp01 returns vector inside 0, 1
func (a Vec2) Clamp01() Vec2 {
	return Vec2{X: Clamp01(a.X), Y: Clamp01(a.Y)}
}

// Clamp1 returns vector inside -1, 1
func (a Vec2) Clamp1() Vec2 {
	return Vec2{X: Clamp1(a.X), Y: Clamp1(a.Y)}
}

// ClampUnit returns scaled vector where Len <= 1
func (a Vec2) ClampUnit() Vec2 {
	m := a.Len()
	if m < 1 {
		return a
	}
	return a.Scale(1 / m)
}

// Slice returns all components as a slice
func (a Vec2) Slice() []float32 {
	return []float32{a.X, a.Y}
}

// Float2 returns two first components
func (a Vec2) FloatXY() [2]float32 { return [2]float32{a.X, a.Y} }

// XY returns Vec2{X, Y}
func (a Vec2) XY() Vec2 { return Vec2{a.X, a.Y} }

type Box2 struct{ Min, Max Vec2 }

// Add returns Box2 translated by p
func (r Box2) Add(p Vec2) Box2 {
	return Box2{Min: r.Min.Add(p), Max: r.Max.Add(p)}
}

// Canon returns canonical version of r
func (r Box2) Canon() Box2 {
	return Box2{Min: r.Min.Min(r.Max), Max: r.Min.Max(r.Max)}
}

// Contains returns whether p is fully contained in rect
func (r Box2) Contains(p Vec2) bool {
	return r.Min.Lt(p) && p.Lt(r.Max)
}

// Inflate increases size by 2*p
func (r Box2) Inflate(p Vec2) Box2 {
	return Box2{Min: r.Min.Sub(p), Max: r.Max.Add(p)}
}

// Deflate shrinks size by 2*p
func (r Box2) Deflate(p Vec2) Box2 {
	return Box2{Min: r.Min.Add(p), Max: r.Max.Sub(p)}
}

// Union returns smallest bounding box that contains both a and b
func (a Box2) Union(b Box2) Box2 {
	a, b = a.Canon(), b.Canon()
	return Box2{Min: a.Min.Min(b.Min), Max: a.Max.Max(b.Max)}
}

// Size returns Max - Min
func (r Box2) Size() Vec2 { return r.Max.Sub(r.Min) }

// Vec3 is a 3 element vector
type Vec3 struct{ X, Y, Z float32 }

var (
	Z3 = Vec3{}
	I3 = Vec3{1, 1, 1}
)

func V3(X, Y, Z float32) Vec3 {
	return Vec3{X: X, Y: Y, Z: Z}
}

// Add does a component wise a + b
func (a Vec3) Add(b Vec3) Vec3 {
	return Vec3{X: a.X + b.X, Y: a.Y + b.Y, Z: a.Z + b.Z}
}

// Sub does a component wise a - b
func (a Vec3) Sub(b Vec3) Vec3 {
	return Vec3{X: a.X - b.X, Y: a.Y - b.Y, Z: a.Z - b.Z}
}

// Mul does a component wise a * b
func (a Vec3) Mul(b Vec3) Vec3 {
	return Vec3{X: a.X * b.X, Y: a.Y * b.Y, Z: a.Z * b.Z}
}

// Div does a component wise a / b
func (a Vec3) Div(b Vec3) Vec3 {
	return Vec3{X: a.X / b.X, Y: a.Y / b.Y, Z: a.Z / b.Z}
}

// Scale scales each component by s
func (a Vec3) Scale(s float32) Vec3 {
	return Vec3{X: a.X * s, Y: a.Y * s, Z: a.Z * s}
}

// Dot returns the dot product of a and b
func (a Vec3) Dot(b Vec3) float32 {
	return (a.X*b.X + a.Y*b.Y + a.Z*b.Z)
}

// Len returns vector length
func (a Vec3) Len() float32 {
	return Sqrt(a.X*a.X + a.Y*a.Y + a.Z*a.Z)
}

// Len2 returns vector length squared
func (a Vec3) Len2() float32 {
	return (a.X*a.X + a.Y*a.Y + a.Z*a.Z)
}

// Normalize returns normalized vector
func (a Vec3) Normalize() Vec3 {
	s := 1.0 / a.Len()
	return Vec3{X: a.X * s, Y: a.Y * s, Z: a.Z * s}
}

// Min does an element wise min
func (a Vec3) Min(b Vec3) Vec3 {
	return Vec3{X: Min(a.X, b.X), Y: Min(a.Y, b.Y), Z: Min(a.Z, b.Z)}
}

// Max does an element wise max
func (a Vec3) Max(b Vec3) Vec3 {
	return Vec3{X: Max(a.X, b.X), Y: Max(a.Y, b.Y), Z: Max(a.Z, b.Z)}
}

// Eq returns whether a == b
func (a Vec3) Eq(b Vec3) bool {
	return (a.X == b.X && a.Y == b.Y && a.Z == b.Z)
}

// Lt returns whether a < b
func (a Vec3) Lt(b Vec3) bool {
	return (a.X < b.X && a.Y < b.Y && a.Z < b.Z)
}

// Le returns whether a <= b
func (a Vec3) Le(b Vec3) bool {
	return (a.X <= b.X && a.Y <= b.Y && a.Z <= b.Z)
}

// Gt returns whether a > b
func (a Vec3) Gt(b Vec3) bool {
	return (a.X > b.X && a.Y > b.Y && a.Z > b.Z)
}

// Ge returns whether a >= b
func (a Vec3) Ge(b Vec3) bool {
	return (a.X >= b.X && a.Y >= b.Y && a.Z >= b.Z)
}

// EqAlmost returns whether Abs(a - b) < eps
func (a Vec3) EqAlmost(b Vec3, eps float32) bool {
	return (Abs(a.X-b.X) < eps && Abs(a.Y-b.Y) < eps && Abs(a.Z-b.Z) < eps)
}

// Lerp returns linearly interpolated vector
func (a Vec3) Lerp(b Vec3, p float32) Vec3 {
	return Vec3{X: Lerp(a.X, b.X, p), Y: Lerp(a.Y, b.Y, p), Z: Lerp(a.Z, b.Z, p)}
}

// Lerp returns linearly interpolated vector
func (a Vec3) LerpClamp(b Vec3, p float32) Vec3 {
	if p < 0 {
		return a
	} else if p > 1 {
		return b
	}
	return Vec3{X: Lerp(a.X, b.X, p), Y: Lerp(a.Y, b.Y, p), Z: Lerp(a.Z, b.Z, p)}
}

// Clamp returns vector inside min, max
func (a Vec3) Clamp(min, max Vec3) Vec3 {
	return Vec3{X: Clamp(a.X, min.X, max.X), Y: Clamp(a.Y, min.Y, max.Y), Z: Clamp(a.Z, min.Z, max.Z)}
}

// Clamp01 returns vector inside 0, 1
func (a Vec3) Clamp01() Vec3 {
	return Vec3{X: Clamp01(a.X), Y: Clamp01(a.Y), Z: Clamp01(a.Z)}
}

// Clamp1 returns vector inside -1, 1
func (a Vec3) Clamp1() Vec3 {
	return Vec3{X: Clamp1(a.X), Y: Clamp1(a.Y), Z: Clamp1(a.Z)}
}

// ClampUnit returns scaled vector where Len <= 1
func (a Vec3) ClampUnit() Vec3 {
	m := a.Len()
	if m < 1 {
		return a
	}
	return a.Scale(1 / m)
}

// Slice returns all components as a slice
func (a Vec3) Slice() []float32 {
	return []float32{a.X, a.Y, a.Z}
}

// Float2 returns two first components
func (a Vec3) FloatXY() [2]float32 { return [2]float32{a.X, a.Y} }

// Float3 returns three first components
func (a Vec3) FloatXYZ() [3]float32 { return [3]float32{a.X, a.Y, a.Z} }

// XY returns Vec2{X, Y}
func (a Vec3) XY() Vec2 { return Vec2{a.X, a.Y} }

// XYZ returns Vec3{X, Y, Z}
func (a Vec3) XYZ() Vec3 { return Vec3{a.X, a.Y, a.Z} }

type Box3 struct{ Min, Max Vec3 }

// Add returns Box3 translated by p
func (r Box3) Add(p Vec3) Box3 {
	return Box3{Min: r.Min.Add(p), Max: r.Max.Add(p)}
}

// Canon returns canonical version of r
func (r Box3) Canon() Box3 {
	return Box3{Min: r.Min.Min(r.Max), Max: r.Min.Max(r.Max)}
}

// Contains returns whether p is fully contained in rect
func (r Box3) Contains(p Vec3) bool {
	return r.Min.Lt(p) && p.Lt(r.Max)
}

// Inflate increases size by 2*p
func (r Box3) Inflate(p Vec3) Box3 {
	return Box3{Min: r.Min.Sub(p), Max: r.Max.Add(p)}
}

// Deflate shrinks size by 2*p
func (r Box3) Deflate(p Vec3) Box3 {
	return Box3{Min: r.Min.Add(p), Max: r.Max.Sub(p)}
}

// Union returns smallest bounding box that contains both a and b
func (a Box3) Union(b Box3) Box3 {
	a, b = a.Canon(), b.Canon()
	return Box3{Min: a.Min.Min(b.Min), Max: a.Max.Max(b.Max)}
}

// Size returns Max - Min
func (r Box3) Size() Vec3 { return r.Max.Sub(r.Min) }

// Vec4 is a 4 element vector
type Vec4 struct{ X, Y, Z, W float32 }

var (
	Z4 = Vec4{}
	I4 = Vec4{1, 1, 1, 1}
)

func V4(X, Y, Z, W float32) Vec4 {
	return Vec4{X: X, Y: Y, Z: Z, W: W}
}

// Add does a component wise a + b
func (a Vec4) Add(b Vec4) Vec4 {
	return Vec4{X: a.X + b.X, Y: a.Y + b.Y, Z: a.Z + b.Z, W: a.W + b.W}
}

// Sub does a component wise a - b
func (a Vec4) Sub(b Vec4) Vec4 {
	return Vec4{X: a.X - b.X, Y: a.Y - b.Y, Z: a.Z - b.Z, W: a.W - b.W}
}

// Mul does a component wise a * b
func (a Vec4) Mul(b Vec4) Vec4 {
	return Vec4{X: a.X * b.X, Y: a.Y * b.Y, Z: a.Z * b.Z, W: a.W * b.W}
}

// Div does a component wise a / b
func (a Vec4) Div(b Vec4) Vec4 {
	return Vec4{X: a.X / b.X, Y: a.Y / b.Y, Z: a.Z / b.Z, W: a.W / b.W}
}

// Scale scales each component by s
func (a Vec4) Scale(s float32) Vec4 {
	return Vec4{X: a.X * s, Y: a.Y * s, Z: a.Z * s, W: a.W * s}
}

// Dot returns the dot product of a and b
func (a Vec4) Dot(b Vec4) float32 {
	return (a.X*b.X + a.Y*b.Y + a.Z*b.Z + a.W*b.W)
}

// Len returns vector length
func (a Vec4) Len() float32 {
	return Sqrt(a.X*a.X + a.Y*a.Y + a.Z*a.Z + a.W*a.W)
}

// Len2 returns vector length squared
func (a Vec4) Len2() float32 {
	return (a.X*a.X + a.Y*a.Y + a.Z*a.Z + a.W*a.W)
}

// Normalize returns normalized vector
func (a Vec4) Normalize() Vec4 {
	s := 1.0 / a.Len()
	return Vec4{X: a.X * s, Y: a.Y * s, Z: a.Z * s, W: a.W * s}
}

// Min does an element wise min
func (a Vec4) Min(b Vec4) Vec4 {
	return Vec4{X: Min(a.X, b.X), Y: Min(a.Y, b.Y), Z: Min(a.Z, b.Z), W: Min(a.W, b.W)}
}

// Max does an element wise max
func (a Vec4) Max(b Vec4) Vec4 {
	return Vec4{X: Max(a.X, b.X), Y: Max(a.Y, b.Y), Z: Max(a.Z, b.Z), W: Max(a.W, b.W)}
}

// Eq returns whether a == b
func (a Vec4) Eq(b Vec4) bool {
	return (a.X == b.X && a.Y == b.Y && a.Z == b.Z && a.W == b.W)
}

// Lt returns whether a < b
func (a Vec4) Lt(b Vec4) bool {
	return (a.X < b.X && a.Y < b.Y && a.Z < b.Z && a.W < b.W)
}

// Le returns whether a <= b
func (a Vec4) Le(b Vec4) bool {
	return (a.X <= b.X && a.Y <= b.Y && a.Z <= b.Z && a.W <= b.W)
}

// Gt returns whether a > b
func (a Vec4) Gt(b Vec4) bool {
	return (a.X > b.X && a.Y > b.Y && a.Z > b.Z && a.W > b.W)
}

// Ge returns whether a >= b
func (a Vec4) Ge(b Vec4) bool {
	return (a.X >= b.X && a.Y >= b.Y && a.Z >= b.Z && a.W >= b.W)
}

// EqAlmost returns whether Abs(a - b) < eps
func (a Vec4) EqAlmost(b Vec4, eps float32) bool {
	return (Abs(a.X-b.X) < eps && Abs(a.Y-b.Y) < eps && Abs(a.Z-b.Z) < eps && Abs(a.W-b.W) < eps)
}

// Lerp returns linearly interpolated vector
func (a Vec4) Lerp(b Vec4, p float32) Vec4 {
	return Vec4{X: Lerp(a.X, b.X, p), Y: Lerp(a.Y, b.Y, p), Z: Lerp(a.Z, b.Z, p), W: Lerp(a.W, b.W, p)}
}

// Lerp returns linearly interpolated vector
func (a Vec4) LerpClamp(b Vec4, p float32) Vec4 {
	if p < 0 {
		return a
	} else if p > 1 {
		return b
	}
	return Vec4{X: Lerp(a.X, b.X, p), Y: Lerp(a.Y, b.Y, p), Z: Lerp(a.Z, b.Z, p), W: Lerp(a.W, b.W, p)}
}

// Clamp returns vector inside min, max
func (a Vec4) Clamp(min, max Vec4) Vec4 {
	return Vec4{X: Clamp(a.X, min.X, max.X), Y: Clamp(a.Y, min.Y, max.Y), Z: Clamp(a.Z, min.Z, max.Z), W: Clamp(a.W, min.W, max.W)}
}

// Clamp01 returns vector inside 0, 1
func (a Vec4) Clamp01() Vec4 {
	return Vec4{X: Clamp01(a.X), Y: Clamp01(a.Y), Z: Clamp01(a.Z), W: Clamp01(a.W)}
}

// Clamp1 returns vector inside -1, 1
func (a Vec4) Clamp1() Vec4 {
	return Vec4{X: Clamp1(a.X), Y: Clamp1(a.Y), Z: Clamp1(a.Z), W: Clamp1(a.W)}
}

// ClampUnit returns scaled vector where Len <= 1
func (a Vec4) ClampUnit() Vec4 {
	m := a.Len()
	if m < 1 {
		return a
	}
	return a.Scale(1 / m)
}

// Slice returns all components as a slice
func (a Vec4) Slice() []float32 {
	return []float32{a.X, a.Y, a.Z, a.W}
}

// Float2 returns two first components
func (a Vec4) FloatXY() [2]float32 { return [2]float32{a.X, a.Y} }

// Float3 returns three first components
func (a Vec4) FloatXYZ() [3]float32 { return [3]float32{a.X, a.Y, a.Z} }

// Float4 returns four first components
func (a Vec4) FloatXYZW() [4]float32 { return [4]float32{a.X, a.Y, a.Z, a.W} }

// XY returns Vec2{X, Y}
func (a Vec4) XY() Vec2 { return Vec2{a.X, a.Y} }

// XYZ returns Vec3{X, Y, Z}
func (a Vec4) XYZ() Vec3 { return Vec3{a.X, a.Y, a.Z} }

// XYZW returns Vec4{X, Y, Z, W}
func (a Vec4) XYZW() Vec4 { return Vec4{a.X, a.Y, a.Z, a.W} }
