.PHONY : all clean

math2017 = day20-v day20-p day20-coll
math2019 = day16-p1 day16-p2 day16-insight day18 day22-ops day22-comp day22-param
math2021 = day01 day06 day06-mat day07

png2017 = $(foreach n,$(math2017),2017-notes-$(n).png)
png2019 = $(foreach n,$(math2019),2019-notes-$(n).png)
png2021 = $(foreach n,$(math2021),2021-notes-$(n).png)

all: $(png2017) $(png2019) $(png2021)

$(png2017) &: ../2017-notes.md genmath
	./genmath $<

$(png2019) &: ../2019-notes.md genmath
	./genmath $<

$(png2021) &: ../2021-notes.md genmath
	./genmath $<

genmath: ../../cmd/genmath/main.go
	go build ../../cmd/genmath

clean:
	$(RM) $(png2017) $(png2019) $(png2021) genmath
