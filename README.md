# LDCOM
## Convert Ludum Dare comments
This programm exports comments on the ludumdare.com page to markdown. The included script `topdf.sh` uses pandoc to also convert to pdf.

```bash
# generate pdf for ludum dare #37 for the game with uuid 26539
./topdf.sh 37 26539
```

You can find your uuid in the link to your entry on ludumdare.com

## Binary Arguments

```
ldcom [-num <num_of_ld>] [-id <id_of_entry>]
```
