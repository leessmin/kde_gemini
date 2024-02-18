# Maintainer: leessmin <1442772970@qq.com>

pkgname=kde_gemini
pkgver=0.0.1
pkgrel=1
pkgdesc="自动切换kde桌面主题"
arch=('x86_64')
url="www.github.com/leessmin/kde_gemini"
license=('LGPL3')
depends=('plasma-desktop>=5.20.0')
options=('!strip')
source=(kde_gemini
		kde_gemini.desktop
        kde_gemini.png)
sha256sums=('8285332102113ccfc9e6029240b19554533ed67d6f170cb6441f864a40cf1fed'
			'9a72f9fb0080e68045ea9aa070d316e5ed6b3fc339e7e1a9d6068c2a6c097656'
			'3125d90dc334b0f4b33269afe73a2c33e7a7f11f92314470d2563e7616d68d02'
			)

package() {
	install -dm755 "${pkgdir}"/usr/bin/
	install -Dm644 "${srcdir}"/${pkgname}.desktop "${pkgdir}"/usr/share/applications/${pkgname}.desktop
	install -Dm644 "${srcdir}"/${pkgname}.png "${pkgdir}"/usr/share/pixmaps/${pkgname}.png
	install -Dm755 "${srcdir}"/${pkgname} "${pkgdir}"/opt/${pkgname}/${pkgname}

	ln -s /opt/${pkgname}/${pkgname} "${pkgdir}"/usr/bin/${pkgname}
}

