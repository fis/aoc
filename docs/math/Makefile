.PHONY : all clean

png2017 = 2017-notes-day20-v.png 2017-notes-day20-p.png 2017-notes-day20-coll.png

all: $(png2017)

$(png2017) &: ../2017-notes.md genmath
	./genmath $<

genmath: ../../cmd/genmath/main.go
	go build ../../cmd/genmath

clean:
	$(RM) $(png2017) genmath
