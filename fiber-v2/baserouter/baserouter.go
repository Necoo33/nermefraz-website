package baserouter

import (
	"github.com/gofiber/fiber/v2"
)

func BaseRouter(server *fiber.App) {
	routes := server.Group("/")

	routes.Get("/", func(context *fiber.Ctx) error {
		return context.Render("index", fiber.Map{
			"Title":       "Nermefraz | Geleceğe Yumuşak Dokunuş",
			"Description": "Nermefraz Teknoloji ile, Website, Yazılım, Digital Marketing, Teknoloji Ve diğer teknoloji alanlarında en ileriye!",
			"Layer":       1,
		})
	})

	routes.Get("/biz", func(context *fiber.Ctx) error {
		return context.Render("about-us", fiber.Map{
			"Title":       "Hakkımzıda | Nermefraz | Geleceğe Yumuşak Dokunuş",
			"Description": "Nermefraz Bilişim Teknolojileri, Yazılım Ve Danışmanlık Şirketi Hakkında Temel Bilgiler",
			"Layer":       1,
		})
	})

	routes.Get("/hizmetlerimiz", func(context *fiber.Ctx) error {
		return context.Render("services", fiber.Map{
			"Title":       "Hizmetlerimiz | Nermefraz | Geleceğe Yumuşak Dokunuş",
			"Description": "Nermefraz Teknoloji'nin estetik ve yüksek animasyonlu web siteler, özel sistemler, kurumsal ve e-ticaret siteler, teknik danışmanlık ve dijital reklamcılık hizmetleriyle ilgili detayları içeren sayfası.",
			"Layer":       1,
		})
	})

	routes.Get("/hizmetlerimiz/website", func(context *fiber.Ctx) error {
		return context.Render("website", fiber.Map{
			"Title":       "Başlangıç Website | Nermefraz | Geleceğe Yumuşak Dokunuş",
			"Description": "Nermefraz Teknoloji'nin estetik görünüm, yüksek google görünürlüğü ve yüksek düzenlenebilirlik vaat ettiği Website hizmeti ile markanızın prestij / itibarını üst seviyeye taşıyın!",
			"Layer":       2,
		})
	})

	routes.Get("/hizmetlerimiz/ozel-yazilim", func(context *fiber.Ctx) error {
		return context.Render("special-software", fiber.Map{
			"Title":       "Özel Yazılım | Nermefraz | Geleceğe Yumuşak Dokunuş",
			"Description": "Nermefraz Teknoloji'nin yüksek performans, düşük kaynak tüketimi, yüksek fonksiyonellik içeren özel yazılım sistemleri oluşturma ve bakımını yapmayı vaat ettiği Özel Yazılım hizmeti ile en uygun çözümleri elde edin!",
			"Layer":       2,
		})
	})

	routes.Get("/hizmetlerimiz/digital-marketing", func(context *fiber.Ctx) error {
		return context.Render("digital-marketing", fiber.Map{
			"Title":       "Digital Marketing | Nermefraz | Geleceğe Yumuşak Dokunuş",
			"Description": "Nermefraz Teknoloji'nin seo optimizasyonu, google desteği ve dijital reklamları içeren sitenizi dijital dünyada görünür kılmayı vaat eden Digital Marketing hizmeti ile kendinizi dünyaya tanıtın!",
			"Layer":       2,
		})
	})

	routes.Get("/hizmetlerimiz/teknik-danismanlik", func(context *fiber.Ctx) error {
		return context.Render("technical-consultancy", fiber.Map{
			"Title":       "Teknik Danışmanlık | Nermefraz | Geleceğe Yumuşak Dokunuş",
			"Description": "Nermefraz Teknoloji'nin siber güvenlik, site optimizasyonu ve erişilebilirlik alanlarında raporlar sunma ve düzenlemeler yapmayı içeren Teknik Danışmanlık hizmeti ile potansiyelinizin ötesine sert iniş yapın!",
			"Layer":       2,
		})
	})

	routes.Get("/projeler", func(context *fiber.Ctx) error {
		return context.Render("projects", fiber.Map{
			"Title":       "Projelerimiz | Nermefraz | Geleceğe Yumuşak Dokunuş",
			"Description": "Nermefraz Teknoloji'nin hayata geçirdiği projelerin sergilendiği, yazılımcılardan son kullanıcıya kadar, her müşteri profiline hitap eden, en başarılı yazılım projeleriyle alakalı detayların yer aldığı sayfa.",
			"Layer":       1,
		})
	})
}
