首先是自己通过继承，对基类同一函数进行不同实现：

#include <iostream>
using namespace std;

class Shape
{
  public:
    void setWidth(int w)
    {
        width = w;
    }
    void setHeight(int h)
    {
        height = h;
    }
    int getArea()
    {
        return 0;
    }
  protected:
    int width;
    int height;
};

class Rectangle:public Shape
{
  public:
    int getArea()
    {
       return (width*height);
    }
};

class Triangle:public Shape
{
    public:
    int getArea()
    {
        return (width*height)/2;
    }

};

int main()
{
    Rectangle rectObj;
    Triangle  triObj;
    
   
    rectObj.setWidth(5);
    rectObj.setHeight(6);
    cout << "Rectangle:"<<rectObj.getArea()<<endl;
    
    triObj.setWidth(5);
    triObj.setHeight(6);
    cout << "Triangle:"<<triObj.getArea()<<endl;
   
    return 0;
}

-------
使用virtual方式：
#include <iostream>
using namespace std;

class Shape
{
  public:
    void setWidth(int w)
    {
        width = w;
    }
    void setHeight(int h)
    {
        height = h;
    }
    // 提供接口框架的纯虚函数
    virtual int getArea()=0;
    
  protected:
    int width;
    int height;
};

//派生类
class Rectangle:public Shape
{
  public:
    int getArea()
    {
       return (width*height);
    }
};

class Triangle:public Shape
{
    public:
    int getArea()
    {
        return (width*height)/2;
    }

};

int main()
{
    Rectangle rectObj;
    Triangle  triObj;
    
    // 输出对象的面积
    rectObj.setWidth(5);
    rectObj.setHeight(6);
    cout << "Rectangle:"<<rectObj.getArea()<<endl;
    // 输出对象的面积
    triObj.setWidth(5);
    triObj.setHeight(6);
    cout << "Triangle:"<<triObj.getArea()<<endl;
   
    return 0;
}
