#include <iostream>
#include <cstdlib> 
/**
 * this pointer points to the object used to  invoke a member function
 * Basically, this is passed as a hidden argument to the method
 * 
 * all class methods have this pointer set to the address of the object that invokes the method
 * you use the -> operator to access structure members via a pointer
*/
using namespace std;
class Box
{
public:
	//constructor
	Box(int l,int w,int h)
	{
		length = l;
		width = w;
		height = h;
	}
	int calVolume()
	{
		return length*width*height;
	}
	//
	bool compareVolume(Box box)
	{
	    return this->calVolume() > box.calVolume() ? 1 : 0 ;
	}
private:
	int length;
	int width;
	int height;
};
// main function
int main( )
{
   Box box1(3,4,5);
   Box box2(2,3,3);
   cout << "the valume of Box：" << box1.calVolume() << endl;
   cout << "the valume of Box：" << box2.calVolume() << endl;
   
   if(box1.compareVolume(box2))
   {
       cout << "box1 is larger than box2" <<endl;
   }else
   {
       cout << "box1 is smaller or equal than box2" <<endl;
   }
   
   return 0;
}
