test query:

 blastp -db dicty_primary_protein -query test_query.fsa -evalue 0.1 -num_alignments 50 -word_size 3 -seg 'yes'

make the db:

makeblastdb -in dicty_primary_protein -dbtype prot
