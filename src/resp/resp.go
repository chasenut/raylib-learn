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
		Spacing: 4,
	}
}

type Button struct {
	Rect rl.Rectangle
	ColorEnabled rl.Color
	ColorDisabled rl.Color
	Text Text
	Toogle bool
	Active bool
}

func NewButton(rect rl.Rectangle,
	colorEnabled, colorDisabled rl.Color,
	text Text,
	toogle, active bool) Button {
	return Button{
		Rect: rect,
		ColorEnabled: colorEnabled,
		ColorDisabled: colorDisabled,
		Text: text,
		Toogle: toogle,
		Active: active,
	}
}

func (b *Button) SetText(text Text) {
	b.Text = text
}

func (b *Button) SetTextContent(text string) {
	b.Text.Content = text
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
	PresetTextEmpty Text = NewTextSimple("", 0, 0, 0, rl.White)
	PresetTextHelloWorld Text = NewTextSimple("Hello, world", 0, 0, 20, rl.White)
)

func (b Button) Draw() {
	if b.Active {
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

