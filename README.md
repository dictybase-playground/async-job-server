test query:

 blastp -db dicty_primary_protein -query test_query.fsa -evalue 0.1 -num_alignments 50 -word_size 3 -seg 'yes'

make the db:

makeblastdb -in dicty_primary_protein -dbtype prot

Requirements:
1. Blastp installed https://blast.ncbi.nlm.nih.gov/Blast.cgi?CMD=Web&PAGE_TYPE=BlastDocs&DOC_TYPE=Download
2. Download database http://dictybase.org/db/cgi-bin/dictyBase/download/blast_databases.pl
3. Create database with
 ```makeblastdb -in dicty_primary_protein -dbtype prot```
