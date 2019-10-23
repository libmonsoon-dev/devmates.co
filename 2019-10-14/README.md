### Problem of the Day:
Given a string, reduce the string by removing 3 or more consecutive identical characters. You should greedily remove characters from left to right.

### Example:
```text
Input: "aaabbbc"
Output: "c"
Why?:
1. Remove 3 'a': "aaabbbc" => "bbbc"
2. Remove 3 'b': "bbbc" => "c"

Input: "aabbbacd"
Output: "cd"
Why?:
1. Remove 3 'b': "aabbbacd" => "aaacd"
2. Remove 3 'a': "aaacd" => "cd"

Input: "aabbccddeeedcba"
Output: ""
Why?:
1. Remove 3 'e': "aabbccddeeedcba" => "aabbccdddcba"
2. Remove 3 'd': "aabbccdddcba" => "aabbcccba"
3. Remove 3 'c': "aabbcccba" => "aabbba"
4. Remove 3 'b': "aabbba" => "aaa"
5. Remove 3 'a': "aaa" => ""

Input: "aabbccddeeedcba"
Output: ""
Why?:
1. Remove 3 'e': "aabbccddeeedcba" => "aabbccdddcba"
2. Remove 3 'd': "aabbccdddcba" => "aabbcccba"
3. Remove 3 'c': "aabbcccba" => "aabbba"
4. Remove 3 'b': "aabbba" => "aaa"
5. Remove 3 'a': "aaa" => ""
```
