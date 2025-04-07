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

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

//leetcode 1 两数之和
//https://leetcode.cn/problems/two-sum/description/
vector<int> find_target_sum_index(vector<int> srcList, int target) {
    vector<int> result;
    //为了减少时间复杂度，以空间换时间即可
    unordered_map<int, int> hashMap;
    //循环遍历整个数组
    int length = srcList.size();
    for (int i=0; i<length; i++) {
        //利用hash表来查询目标值
        auto f = hashMap.find(target - srcList[i]);
        if (f != hashMap.end()) {
            result.push_back(f->second);
            result.push_back(i);
            return result;
        }
        hashMap[srcList[i]]= i;
    }

    return result;
}

void test_find_target_sum_index() {
    vector<int> src;
    src.push_back(1);
    src.push_back(2);
    src.push_back(3);
    src.push_back(4);

    vector<int> res = find_target_sum_index(src, 3);
    if (!res.empty()) {
        std::cout<<res[0]<<"-"<<res[1]<<std::endl;
    }
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//单链表翻转
struct listNode {
    int val;
    listNode * next;
};

listNode* reverse_list_node(listNode* head) {
    if (head == NULL || head->next == NULL) {
        return head;
    }

    listNode * cur = head;
    listNode * re = NULL;

    while (cur != NULL) {
        listNode * temp = cur;
        cur = cur->next;
        temp->next = re;
        re = temp;
    }

    return re;
}
//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//求树深度
struct dep_binary_tree {
    int value;
    dep_binary_tree * left;
    dep_binary_tree * right;
};

//递归解法
int get_binary_max_length(dep_binary_tree * root) {
    if (root == NULL) {
        return 0;
    }
   return max(get_binary_max_length(root->left), get_binary_max_length(root->right)) + 1;
}



//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
int main() {
    test_find_target_sum_index();
    return 0;
}
