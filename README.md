# Flesch Index Checker

## Summary

The purpose of this tool is to demonstrate my ability to process data for a graduate level course in machine learning. The Flesch Index is a score that indicates approximate textual complexity of a passage based upon a fairly simple formula. The language of machine learning is Python, but as an experienced developer I have chosen to explore another alternative to broaden my horizons. Therefore, this project uses Go to calculate and produce visualizations that describe readability characteristics of documents.

## Execution

In this case, the test material will be provided by public sources available from The Gutenburg Project.
- NY Times article on Health Care
- Lincoln's Gettysburg Address
- Moby Dick

### Formula

Flesch_Index = 206.835−84.6 (𝑛𝑢𝑚𝑆𝑦𝑙𝑙𝑎𝑏𝑙𝑒𝑠 / 𝑛𝑢𝑚𝑊𝑜𝑟𝑑𝑠)−1.015 (𝑛𝑢𝑚𝑊𝑜𝑟𝑑𝑠 / 𝑛𝑢𝑚𝑆𝑒𝑛𝑡𝑒𝑛𝑐𝑒𝑠)

### Sentence

Consider a sentence to have been encountered whenever you find a word that ends in a specific punctuation symbol: a period, question mark, or exclamation point.

### Word

A word is a contiguous sequence of alphabetic characters.  Whitespace defines word boundaries.

### Syllable

A syllable is considered to have been encountered whenever you detect:
- Rule 1:  a vowel at the start of a wordor
- Rule 2:  a vowel following a consonant in a word
One exception to Rule 2: a lone ‘e’ at the end of a word does not count as a syllable.

## Practical Application and Limitations

I don't anticipate particularly valuable usage of this library for others. The intent is to demonstrate to my professor that I can create reasonable applications that process data, and as a software engineer I had better deliver, even if this language varies from my daily driver (C# and Typescript). Go is a personal favorite, so I intent to learn about and demonstrate the use of gonum especially.

## License and Usage Guidelines

I have licensed this as under an MIT license, which provides broad rights for reuse. However, if you are a student somewhere with a similar project, your ethical codes apply and you might not be allowed to use everything you see here based upon your academic guidelines.

I approached my professor, Dr. Wolffe, for permission to host this in a public repo and permission was granted.