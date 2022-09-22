// Code generated by qtc from "base.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// 这是我们模板的基础实现

//line templates/admin/base.qtpl:2
package admin

//line templates/admin/base.qtpl:2
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line templates/admin/base.qtpl:2
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line templates/admin/base.qtpl:3
type Partial interface {
//line templates/admin/base.qtpl:3
	Body() string
//line templates/admin/base.qtpl:3
	StreamBody(qw422016 *qt422016.Writer)
//line templates/admin/base.qtpl:3
	WriteBody(qq422016 qtio422016.Writer)
//line templates/admin/base.qtpl:3
}

// 模板编写实现Partial接口的模板

//line templates/admin/base.qtpl:8
func StreamTemplate(qw422016 *qt422016.Writer, p Partial) {
//line templates/admin/base.qtpl:8
	qw422016.N().S(`
            `)
//line templates/admin/base.qtpl:9
	p.StreamBody(qw422016)
//line templates/admin/base.qtpl:9
	qw422016.N().S(`
`)
//line templates/admin/base.qtpl:10
}

//line templates/admin/base.qtpl:10
func WriteTemplate(qq422016 qtio422016.Writer, p Partial) {
//line templates/admin/base.qtpl:10
	qw422016 := qt422016.AcquireWriter(qq422016)
//line templates/admin/base.qtpl:10
	StreamTemplate(qw422016, p)
//line templates/admin/base.qtpl:10
	qt422016.ReleaseWriter(qw422016)
//line templates/admin/base.qtpl:10
}

//line templates/admin/base.qtpl:10
func Template(p Partial) string {
//line templates/admin/base.qtpl:10
	qb422016 := qt422016.AcquireByteBuffer()
//line templates/admin/base.qtpl:10
	WriteTemplate(qb422016, p)
//line templates/admin/base.qtpl:10
	qs422016 := string(qb422016.B)
//line templates/admin/base.qtpl:10
	qt422016.ReleaseByteBuffer(qb422016)
//line templates/admin/base.qtpl:10
	return qs422016
//line templates/admin/base.qtpl:10
}

// 基本模板实现。 如果需要，其他页面可以继承
// 仅覆盖某些部分方法。

//line templates/admin/base.qtpl:13
type Base struct{}

//line templates/admin/base.qtpl:14
func (b *Base) StreamBody(qw422016 *qt422016.Writer) {
//line templates/admin/base.qtpl:14
	qw422016.N().S(`This is the base body`)
//line templates/admin/base.qtpl:14
}

//line templates/admin/base.qtpl:14
func (b *Base) WriteBody(qq422016 qtio422016.Writer) {
//line templates/admin/base.qtpl:14
	qw422016 := qt422016.AcquireWriter(qq422016)
//line templates/admin/base.qtpl:14
	b.StreamBody(qw422016)
//line templates/admin/base.qtpl:14
	qt422016.ReleaseWriter(qw422016)
//line templates/admin/base.qtpl:14
}

//line templates/admin/base.qtpl:14
func (b *Base) Body() string {
//line templates/admin/base.qtpl:14
	qb422016 := qt422016.AcquireByteBuffer()
//line templates/admin/base.qtpl:14
	b.WriteBody(qb422016)
//line templates/admin/base.qtpl:14
	qs422016 := string(qb422016.B)
//line templates/admin/base.qtpl:14
	qt422016.ReleaseByteBuffer(qb422016)
//line templates/admin/base.qtpl:14
	return qs422016
//line templates/admin/base.qtpl:14
}