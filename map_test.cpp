#include <map>
#include <string>
#include <iostream>
using namespace std;

int main()
{
    map<int,string> mapStudent;
    mapStudent[1] = "student_one";
    mapStudent[2] = "student_Two";
    mapStudent[3] = "student_Three";
    
    //使用反向迭代器 
    map<int,string>::reverse_iterator iter;
    for(iter = mapStudent.rbegin(); iter != mapStudent.rend(); iter++)
        std::cout << iter->first <<" "<<iter->second << std::endl;
        
    //使用数组 （可覆盖 ）
    int iSize = mapStudent.size();
    for(int i = 1; i <= iSize; i++)
        cout << i << mapStudent[i] <<endl;
        
    
    return 0;
}



------
将map按value排序
注：本功能主要是考虑到sort只能对序列容器（线性存储）进行排序，因此这里要对map的value进行排序
    首先把map的元素按pair形式插入到vector中，然后利用sort排序（用一个新写的比较函数），这样可以实现。
    
注：map中first是指key,Second是指value。因此本功能实现根据学生分数增序排列


#include <map>
#include <vector> 
#include <string>
#include <iostream>
#include <algorithm>
using namespace std;

typedef pair<string, int> PAIR;
bool cmp_by_value(const PAIR & lhs, const PAIR& rhs)
{
    return lhs.second < rhs.second;
}
struct CmpBYValue{
    bool operator()(const PAIR& lhs, const PAIR& rhs){
        return lhs.second < rhs.second;
    }
};
int main()
{
    map<string,int> name_score_map;
    name_score_map["LiMin"] = 90;
    name_score_map["ZiLinMi"] = 79;
    name_score_map["BoB"] = 92;
    name_score_map.insert(make_pair("Bing",69));
    name_score_map.insert(make_pair("Albert",86));
    
    //把map中元素转存到vector中  
    vector<PAIR> name_score_vec(name_score_map.begin(),name_score_map.end());
    //sort(name_score_vec.begin(),name_score_vec.end(),CmpBYValue());
    sort(name_score_vec.begin(),name_score_vec.end(),cmp_by_value);
    
    for(int i=0; i!=name_score_vec.size(); ++i){
        cout<< name_score_vec[i].first << " " <<name_score_vec[i].second<<endl;
    }
    
    return 0;
}
