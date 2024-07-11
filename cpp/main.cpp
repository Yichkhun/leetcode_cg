#include <iostream>
#include <unordered_set>

using namespace std;

//leetcode 3 无重复字符串的最长子串
//https://leetcode.cn/problems/longest-substring-without-repeating-characters/description/
int find_max_len_sub_str(const string & str) {
    int length = str.size();
    //定义一个存放不重复字符的数据集合
    unordered_set<char> filterSet;
    //最终存储不重复子串的右侧索引
    int right = 0;
    //最大长度
    int maxLength = 0;
    //循环扫描整个字符串 这里的i相当于最终存储不重复子串的左侧索引
    for(int i=0; i<length; i++) {
        //如果非第一次扫描，可以理解为无重复子串的左索引需要移动， 因为重复的肯定是右侧索引的最后一个和左侧的第一个重复，所以需要移动左侧第一个。
        if (i!=0) {
            filterSet.erase(str[i-1]);
        }
        //如果是一直不重复，那就持续添加到集合中
        while(right < length && !filterSet.count(str[right])) {
            filterSet.insert(str[right]);
            right++;
        }
        //取之前判定认为最大的长度和这轮扫描后 右侧索引值减去左侧索引值的长度
        maxLength = max(maxLength, right - i);
    }
    return maxLength;
}

void test_find_max_len_sub_str() {
    //exp1
    const string s1 = "abcabcbb";
    int max = find_max_len_sub_str(s1);
    std::cout<<"max-len"<<max<<std::endl;

    //exp2
    const string s2 = "bbbbb";
    max = find_max_len_sub_str(s2);
    std::cout<<"max-len"<<max<<std::endl;

    //exp3
    const string s3 = "pwwkew";
    max = find_max_len_sub_str(s3);
    std::cout<<"max-len"<<max<<std::endl;
}

int main() {
    test_find_max_len_sub_str();
    return 0;
}
