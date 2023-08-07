#include <iostream>
#include <map>

using namespace std;

int main()
{
	map<int,int> mp;
	mp.insert(make_pair(1,10));
	mp.insert(make_pair(1,10));

	cout << mp.size() << endl;

	return 0;
}
