package poppler

// #cgo pkg-config: --cflags-only-I poppler-glib
// #include <poppler.h>
// #include <glib.h>
import (
	"C"
)

type Page struct {
	p poppDoc
}

func (p *Page) Text() string {
	return C.GoString(C.poppler_page_get_text(p.p))
}

func (p *Page) TextAttributes() (results []TextAttributes) {
	a := C.poppler_page_get_text_attributes(p.p)
	defer C.poppler_page_free_text_attributes(a)
	var attr *C.PopplerTextAttributes
	results = make([]TextAttributes, 0)
	for el := C.g_list_first(a); el != nil; el = el.next {
		attr = (*C.PopplerTextAttributes)(el.data)
		fn := *attr.font_name
		result := TextAttributes{
			FontName:     toString(&fn),
			FontSize:     float64(attr.font_size),
			IsUnderlined: toBool(attr.is_underlined),
			StartIndex:   int(attr.start_index),
			EndIndex:     int(attr.end_index),
			Color: Color{
				R: int(attr.color.red),
				G: int(attr.color.green),
				B: int(attr.color.blue),
			},
		}
		results = append(results, result)
	}
	return
}

func (p *Page) Size() (width, height float64) {
	var w, h C.double
	C.poppler_page_get_size(p.p, &w, &h)
	return float64(w), float64(h)
}

func (p *Page) Index() int {
	return int(C.poppler_page_get_index(p.p))
}

func (p *Page) Label() string {
	return toString(C.poppler_page_get_label(p.p))
}

func (p *Page) Duration() float64 {
	return float64(C.poppler_page_get_duration(p.p))
}

func (p *Page) Images() (results []Image) {
	l := C.poppler_page_get_image_mapping(p.p)
	defer C.poppler_page_free_image_mapping(l)
	results = make([]Image, 0)
	var im *C.PopplerImageMapping
	for el := C.g_list_first(l); el != nil; el = el.next {
		im = (*C.PopplerImageMapping)(el.data)
		result := Image{
			Id: int(im.image_id),
			Area: Rectangle{
				X1: float64(im.area.x1),
				Y1: float64(im.area.y1),
				X2: float64(im.area.x2),
				Y2: float64(im.area.y2),
			},
			p: p.p,
		}
		results = append(results, result)
	}
	return
}

func (p *Page) TextLayout() (layouts []Rectangle) {
	
	return
}