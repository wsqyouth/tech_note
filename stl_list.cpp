list 是可逆（双向）容器，同时又是顺序容器

► 特点
□ 在任意位置插入和删除元素都很快
□ 不支持随机访问
► 接合(splice)操作
□ s1.splice(p, s2, q1, q2)：将s2中[q1,q2)移动到s1中p所指向元素之前


-----
#include <iostream>
#include <list>
using namespace std;

int main()
{
	//默认构造函数构造空容器 
	list<int> coll;
	//填充List 
	for(int i=0;i<=9;i++)
	{
		coll.push_back(i);
	}
	//判断容器元素个数
	cout << coll.size() <<endl; 
	//读取数据
	list<int>::iterator pos = coll.begin();
	while(pos != coll.end())
	{
		cout << *pos <<" ";
		pos++;
	} 
	//容器清空
	coll.clear();
	//判断容器是否为空
	if(coll.empty())
	{
		cout <<"\n容器为空"<<endl;
	}

	
	return 0;
}
