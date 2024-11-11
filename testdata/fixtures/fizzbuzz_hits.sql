TRUNCATE TABLE fizzbuzz_requests RESTART IDENTITY;

INSERT INTO fizzbuzz_requests (int1, int2, max_limit, str1, str2, hits) VALUES
(3, 5, 100, 'fizz', 'buzz', 42),
(2, 4, 50, 'foo', 'bar', 30),
(5, 8, 50, 'john',  'doe', 10);
