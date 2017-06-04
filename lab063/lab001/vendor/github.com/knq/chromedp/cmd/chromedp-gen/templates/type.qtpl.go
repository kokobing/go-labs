// This file is automatically generated by qtc from "type.qtpl".
// See https://github.com/valyala/quicktemplate for details.

//line templates/type.qtpl:1
package templates

//line templates/type.qtpl:1
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line templates/type.qtpl:1
import (
	"strconv"
	"strings"

	"github.com/knq/chromedp/cmd/chromedp-gen/internal"
)

// TypeTemplate is a template for a Typable type.

//line templates/type.qtpl:9
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line templates/type.qtpl:9
func StreamTypeTemplate(qw422016 *qt422016.Writer, t *internal.Type, prefix, suffix string, d *internal.Domain, domains []*internal.Domain, v interface{}, noExposeOverride, omitOnlyWhenOptional bool) {
	//line templates/type.qtpl:10
	typ := prefix + t.CamelName() + suffix

	var extra []*internal.Type
	switch x := v.(type) {
	case []*internal.Type:
		extra = x
	}

	//line templates/type.qtpl:17
	qw422016.N().S(`
`)
	//line templates/type.qtpl:18
	qw422016.N().S(formatComment(t.GetDescription(), "", typ+" "))
	//line templates/type.qtpl:18
	qw422016.N().S(`
type `)
	//line templates/type.qtpl:19
	qw422016.N().S(typ)
	//line templates/type.qtpl:19
	qw422016.N().S(` `)
	//line templates/type.qtpl:19
	qw422016.N().S(t.GoTypeDef(d, domains, extra, noExposeOverride, omitOnlyWhenOptional))
	//line templates/type.qtpl:19
	qw422016.N().S(`
`)
	//line templates/type.qtpl:20
	if t.Parameters == nil && t.Type != internal.TypeArray && t.Type != internal.TypeObject && t.Type != internal.TypeAny {
		//line templates/type.qtpl:21
		gz := t.Type.GoType()
		z := gz
		if strings.Contains(z, ".") {
			z = z[strings.Index(z, ".")+1:]
		}
		z = strings.ToUpper(z[:1]) + z[1:]

		//line templates/type.qtpl:27
		qw422016.N().S(`
// `)
		//line templates/type.qtpl:28
		qw422016.N().S(z)
		//line templates/type.qtpl:28
		qw422016.N().S(` returns the `)
		//line templates/type.qtpl:28
		qw422016.N().S(typ)
		//line templates/type.qtpl:28
		qw422016.N().S(` as `)
		//line templates/type.qtpl:28
		qw422016.N().S(gz)
		//line templates/type.qtpl:28
		qw422016.N().S(` value.
func (t `)
		//line templates/type.qtpl:29
		qw422016.N().S(typ)
		//line templates/type.qtpl:29
		qw422016.N().S(`) `)
		//line templates/type.qtpl:29
		qw422016.N().S(z)
		//line templates/type.qtpl:29
		qw422016.N().S(`() `)
		//line templates/type.qtpl:29
		qw422016.N().S(gz)
		//line templates/type.qtpl:29
		qw422016.N().S(` {
	return `)
		//line templates/type.qtpl:30
		qw422016.N().S(gz)
		//line templates/type.qtpl:30
		qw422016.N().S(`(t)
}
`)
		//line templates/type.qtpl:32
	}
	//line templates/type.qtpl:32
	qw422016.N().S(`
`)
	//line templates/type.qtpl:33
	if ev := t.EnumValues(); ev != nil {
		//line templates/type.qtpl:34
		gz := t.Type.GoType()
		z := gz
		if strings.Contains(z, ".") {
			z = z[strings.Index(z, ".")+1:]
		}
		z = strings.ToUpper(z[:1]) + z[1:]

		//line templates/type.qtpl:40
		qw422016.N().S(`// `)
		//line templates/type.qtpl:40
		qw422016.N().S(typ)
		//line templates/type.qtpl:40
		qw422016.N().S(` values.
const (`)
		//line templates/type.qtpl:41
		for i, e := range ev {
			//line templates/type.qtpl:42
			n := t.EnumValueName(e)
			val := `"` + e + `"`
			if t.Type == internal.TypeInteger && t.EnumBitMask {
				val = strconv.Itoa(1 << uint(i-1))
			} else if t.Type == internal.TypeInteger {
				val = strconv.Itoa(i + 1)
			}

			//line templates/type.qtpl:49
			qw422016.N().S(`
	`)
			//line templates/type.qtpl:50
			qw422016.N().S(n)
			//line templates/type.qtpl:50
			qw422016.N().S(` `)
			//line templates/type.qtpl:50
			qw422016.N().S(typ)
			//line templates/type.qtpl:50
			qw422016.N().S(` = `)
			//line templates/type.qtpl:50
			qw422016.N().S(val)
			//line templates/type.qtpl:50
		}
		//line templates/type.qtpl:50
		qw422016.N().S(`
)
`)
		//line templates/type.qtpl:52
		if t.Type != internal.TypeString {
			//line templates/type.qtpl:52
			qw422016.N().S(`
// String returns the `)
			//line templates/type.qtpl:53
			qw422016.N().S(typ)
			//line templates/type.qtpl:53
			qw422016.N().S(` as string value.
func (t `)
			//line templates/type.qtpl:54
			qw422016.N().S(typ)
			//line templates/type.qtpl:54
			qw422016.N().S(`) String() string {
	switch t {`)
			//line templates/type.qtpl:55
			for _, e := range t.EnumValues() {
				//line templates/type.qtpl:55
				qw422016.N().S(`
	case `)
				//line templates/type.qtpl:56
				qw422016.N().S(t.EnumValueName(e))
				//line templates/type.qtpl:56
				qw422016.N().S(`:
		return "`)
				//line templates/type.qtpl:57
				qw422016.N().S(e)
				//line templates/type.qtpl:57
				qw422016.N().S(`"`)
				//line templates/type.qtpl:57
			}
			//line templates/type.qtpl:57
			qw422016.N().S(`
	}

	return fmt.Sprintf("`)
			//line templates/type.qtpl:60
			qw422016.N().S(typ)
			//line templates/type.qtpl:60
			qw422016.N().S(`(%d)", t)
}
`)
			//line templates/type.qtpl:62
		}
		//line templates/type.qtpl:62
		qw422016.N().S(`

// MarshalEasyJSON satisfies easyjson.Marshaler.
func (t `)
		//line templates/type.qtpl:65
		qw422016.N().S(typ)
		//line templates/type.qtpl:65
		qw422016.N().S(`) MarshalEasyJSON(out *jwriter.Writer) {
	out.`)
		//line templates/type.qtpl:66
		qw422016.N().S(z)
		//line templates/type.qtpl:66
		qw422016.N().S(`(`)
		//line templates/type.qtpl:66
		qw422016.N().S(gz)
		//line templates/type.qtpl:66
		qw422016.N().S(`(t))
}

// MarshalJSON satisfies json.Marshaler.
func (t `)
		//line templates/type.qtpl:70
		qw422016.N().S(typ)
		//line templates/type.qtpl:70
		qw422016.N().S(`) MarshalJSON() ([]byte, error) {
	return easyjson.Marshal(t)
}

// UnmarshalEasyJSON satisfies easyjson.Unmarshaler.
func (t *`)
		//line templates/type.qtpl:75
		qw422016.N().S(typ)
		//line templates/type.qtpl:75
		qw422016.N().S(`) UnmarshalEasyJSON(in *jlexer.Lexer) {
	switch `)
		//line templates/type.qtpl:76
		qw422016.N().S(typ)
		//line templates/type.qtpl:76
		qw422016.N().S(`(in.`)
		//line templates/type.qtpl:76
		qw422016.N().S(z)
		//line templates/type.qtpl:76
		qw422016.N().S(`()) {`)
		//line templates/type.qtpl:76
		for _, e := range t.EnumValues() {
			//line templates/type.qtpl:77
			n := t.EnumValueName(e)

			//line templates/type.qtpl:78
			qw422016.N().S(`
	case `)
			//line templates/type.qtpl:79
			qw422016.N().S(n)
			//line templates/type.qtpl:79
			qw422016.N().S(`:
		*t = `)
			//line templates/type.qtpl:80
			qw422016.N().S(n)
			//line templates/type.qtpl:80
		}
		//line templates/type.qtpl:80
		qw422016.N().S(`

	default:
		in.AddError(errors.New("unknown `)
		//line templates/type.qtpl:83
		qw422016.N().S(typ)
		//line templates/type.qtpl:83
		qw422016.N().S(` value"))
	}
}

// UnmarshalJSON satisfies json.Unmarshaler.
func (t *`)
		//line templates/type.qtpl:88
		qw422016.N().S(typ)
		//line templates/type.qtpl:88
		qw422016.N().S(`) UnmarshalJSON(buf []byte) error {
	return easyjson.Unmarshal(buf, t)
}`)
		//line templates/type.qtpl:90
	}
	//line templates/type.qtpl:90
	qw422016.N().S(`
`)
	//line templates/type.qtpl:91
	if t.Extra != "" {
		//line templates/type.qtpl:91
		qw422016.N().S(`
`)
		//line templates/type.qtpl:92
		qw422016.N().S(t.Extra)
		//line templates/type.qtpl:92
	}
	//line templates/type.qtpl:92
	qw422016.N().S(`
`)
//line templates/type.qtpl:93
}

//line templates/type.qtpl:93
func WriteTypeTemplate(qq422016 qtio422016.Writer, t *internal.Type, prefix, suffix string, d *internal.Domain, domains []*internal.Domain, v interface{}, noExposeOverride, omitOnlyWhenOptional bool) {
	//line templates/type.qtpl:93
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line templates/type.qtpl:93
	StreamTypeTemplate(qw422016, t, prefix, suffix, d, domains, v, noExposeOverride, omitOnlyWhenOptional)
	//line templates/type.qtpl:93
	qt422016.ReleaseWriter(qw422016)
//line templates/type.qtpl:93
}

//line templates/type.qtpl:93
func TypeTemplate(t *internal.Type, prefix, suffix string, d *internal.Domain, domains []*internal.Domain, v interface{}, noExposeOverride, omitOnlyWhenOptional bool) string {
	//line templates/type.qtpl:93
	qb422016 := qt422016.AcquireByteBuffer()
	//line templates/type.qtpl:93
	WriteTypeTemplate(qb422016, t, prefix, suffix, d, domains, v, noExposeOverride, omitOnlyWhenOptional)
	//line templates/type.qtpl:93
	qs422016 := string(qb422016.B)
	//line templates/type.qtpl:93
	qt422016.ReleaseByteBuffer(qb422016)
	//line templates/type.qtpl:93
	return qs422016
//line templates/type.qtpl:93
}
