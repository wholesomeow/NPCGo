COPY ngram_fantasy(id, ngram_value, ngram_posibility)
FROM '/rawdata/csv/Fantasy_Names_NGrams.csv'
DELIMITER ','
CSV HEADER;