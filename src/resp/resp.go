package resp

import rl "github.com/gen2brain/raylib-go/raylib"

type Clickable interface {
	todo()
}

type Text struct {
	Content string
	Pos rl.Vector2
	FontSize float32
	Color rl.Color
	Font rl.Font
	Spacing float32
}

func NewText(content string,
	x, y, fontSize float32,
	color rl.Color,
	font rl.Font,
	spacing float32,
) Text {
	return Text{
		Content: content,
		Pos: rl.NewVector2(x, y),
		FontSize: fontSize,
		Color: color,
		Font: font,
		Spacing: spacing,
	}
}

func NewTextSimple(content string,
	x, y, fontSize float32,
	color rl.Color,
) Text {
	font := rl.GetFontDefault()
	return Text{
		Content: content,
		Pos: rl.NewVector2(x, y),
		FontSize: fontSize,
		Color: color,
		Font: font,
		Spacing: 0,
	}
}

type Button struct {
	active bool
	Toogle bool
	Rect rl.Rectangle
	ColorEnabled rl.Color
	ColorDisabled rl.Color
	Text Text
}

type FixedButton struct {
	Button
	FixedPos rl.Vector2
}

type StaticButton struct {
	Button
}

var (
	DefaultFont rl.Font = rl.GetFontDefault()
)

func (b Button) GetStatus() bool {
	return b.active
}

func (b *Button) SetStatus(status bool) {
	b.active = status
}

func (b Button) Draw() {
	if b.active {
		rl.DrawRectangleRec(b.Rect, b.ColorEnabled)
	} else {
		rl.DrawRectangleRec(b.Rect, b.ColorDisabled)
	}
	textDims := rl.MeasureTextEx(b.Text.Font, b.Text.Content, 
		b.Text.FontSize, b.Text.Spacing)
	rl.DrawTextPro(b.Text.Font, 
		b.Text.Content, 
		b.Text.Pos, 
		rl.NewVector2(textDims.X/2, textDims.Y/2), 
		0, 
		b.Text.FontSize, 
		b.Text.Spacing, 
		b.Text.Color)

}

func (b FixedButton) Update(cam rl.Camera2D) {
	b.Rect.X = b.FixedPos.X + cam.Offset.X
	b.Rect.Y = b.FixedPos.Y + cam.Offset.Y
}

