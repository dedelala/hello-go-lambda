d = $(shell pwd)
p = 1337

.PHONY: dev preso

dev:
	docker run -p $(p):80 \
	-v $(d)/preso.md:/usr/share/nginx/html/slides.md \
	-v $(d)/style.css:/usr/share/nginx/html/styles.css \
	-v $(d):/usr/share/nginx/html/preso \
	-e DEV_MODE=true \
	hairyhenderson/remarkjs

preso:
	docker run -p $(p):80 \
	-v $(d)/preso.md:/slides.md \
	-v $(d)/style.css:/styles.css \
	-v $(d):/usr/share/nginx/html/preso \
	hairyhenderson/remarkjs
