install:
	v .
	mv dyn /usr/bin/dyn
push:
	git add .
	git commit -m "Update By Trusted Contributor"
	git push
