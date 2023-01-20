NAME = sus
all: $(NAME)Doc.pdf
$(NAME)Doc.pdf: $(NAME)Doc.tex $(NAME).tex pdf.ps
	latex $(NAME)Doc
	bibtex $(NAME)Doc
	latex $(NAME)Doc
	latex $(NAME)Doc
	latex $(NAME)Doc
	dvips $(NAME)Doc -o -q
	ps2pdf -dALLOWPSTRANSPARENCY $(NAME)Doc.ps
$(NAME).tex: ../src/$(NAME).org
	bash ../scripts/org2nw ../src/$(NAME).org | awk -f ../scripts/preWeave.awk | noweave -n -x | sed 's/_test/\\_test/' > $(NAME).tex
pdf.ps: pdf.gp pdf.dat
	gnuplot pdf.gp
clean:
	rm -f $(NAME).tex $(NAME)_test.tex *.ps *.pdf *.aux *.bbl *.blg *.dvi *.log *.toc 
