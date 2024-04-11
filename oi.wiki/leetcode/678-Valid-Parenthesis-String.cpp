#include <string>

using std::string;

bool checkValidString(string s) {
    int match = 0, lcnt = 0, rcnt = 0;
    for (auto &&ch : s) {
        switch (ch) {
		case '*':
			lcnt++;
			rcnt = std::min(rcnt + 1, match);
			break;
		case '(':
			match++;
			break;
		case ')':
			if (match > 0) {
				match--;
				rcnt = std::min(rcnt, match);
				break;
			}
			if (lcnt > 0) {
				lcnt--;
				break;
			}
			return false;
		default:
			break;
        }
    }
    return rcnt >= match;
}
