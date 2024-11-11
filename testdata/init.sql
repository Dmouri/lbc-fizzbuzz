CREATE TABLE fizzbuzz_requests (
   int1 INTEGER NOT NULL,
   int2 INTEGER NOT NULL,
   max_limit INTEGER NOT NULL,
   str1 VARCHAR(50) NOT NULL,
   str2 VARCHAR(50) NOT NULL,
   hits INTEGER DEFAULT 1,
   PRIMARY KEY (int1, int2, max_limit, str1, str2)
);
