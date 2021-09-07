int balancedStringSplit(char * s) {
    int ret = 0, cnt = 0;
    for (int i = 0; s[i]; i++) {
		(s[i] == 'L') ? cnt++ : cnt--;
        if (cnt == 0) { ret++; }
    }
    return ret;
}
