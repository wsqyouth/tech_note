#include <iostream>
#include <cstdlib> 
using namespace std;
/*
* environment :64-bit compiler
* identify the difference of sizeof(array/element/pointer)
*/ 
int main(void) { 
	
	int a[10];
	//the same addrss ,but different meanings 
	printf("address:%p\n",a);
	printf("address: %p\n",&a);
	printf("address:%p\n",&a[0]);
	
	//size :the element of arr  
	printf("size: %ld\n",sizeof(&a[12])); //8 
	//size : the array
	printf("size: %ld\n",sizeof(a));    //4*10 
	
	
	
	 int arr[12];
     cout <<"size:"<<sizeof(arr) <<endl;
     int *p = arr;
     cout <<"size:"<<sizeof(p) <<endl;  //8 
     
    int *ptr = (int*)malloc(sizeof(int) * 12);
    //分配12个（可根据实际需要替换该数值）整型存储单元，
    //并将这12个连续的整型存储单元的首地址存储到指针变量p中
    cout <<"size:"<<sizeof(ptr) <<endl;  //8 
    
    free(ptr);
    ptr = NULL;
	return 0;
}
