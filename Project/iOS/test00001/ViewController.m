//
//  ViewController.m
//  test00001
//
//  Created by Murphy on 2018/6/5.
//  Copyright © 2018 Murphy. All rights reserved.
//

#import "ViewController.h"
#import "golib.h"
//#import <Framework/Framework.h>
#import <Api/Api.h>
#import <Framework/Framework.h>

@interface XXModel:NSObject
@property (nonatomic,strong) NSString *name;
@property (nonatomic,assign) NSInteger age;
@property (nonatomic,strong) NSArray *subModels;
@property (nonatomic,strong) XXModel *cModel;
@end

@implementation XXModel

@end


@interface ViewController ()

@end

@implementation ViewController

- (void)viewDidLoad {
    [super viewDidLoad];
    // Do any additional setup after loading the view, typically from a nib.
//    NSString *name1 = @"name_1";
//    NSInteger age1 = 222;
//    NSString *name2 = @"name_2";
//    NSInteger age2 = 3333;
//
//    XXModel *m1 = [[XXModel alloc] init];
//    XXModel *m2 = [[XXModel alloc] init];
//
//    m1.name = name1;
//    m1.age = age1;
//    m2.name = name2;
//    m2.age = age2;
//    name1 = nil;
//    age1 = 0;
//    name2 = nil;
//    age2 = 0;
//    NSArray *arr = [NSArray arrayWithObject:m2];
//    m1.subModels = arr;
//    m1.cModel = m2;
//    m2 = nil;
//    arr = nil;
//    NSLog(@"%@",m1);
//    printf("This is a C Application.\n");
//    GoString name = {(char*)"King", 4};
//    SayHello(name);
//    GoSlice buf = {(void*)"King", 4, 4};
//    SayHelloByte(buf);
//    SayBye();
    NSString *path = NSHomeDirectory();
    BOOL isEx = NO;
    NSError *error = nil;
//    FrameworkIsExist(@"/Users/murphy/Desktop/test1.txt", &isEx, &error);
//    NSLog(@"%@,%@",@(isEx),error);
////
//    error = nil;
//    FrameworkIsExist(@"/Users/murphy/Desktop/test10.txt", &isEx, &error);
//    NSLog(@"%@,%@",@(isEx),error);
//    error = nil;
//    FrameworkIsExist(path, &isEx, &error);
//    NSLog(@"%@,%@",@(isEx),error);
    NSString *name = ApiLogin(@"/Users/murphy/Desktop/test10.txt", &error);
    NSLog(@"%@,%@",name,error);

}


- (void)didReceiveMemoryWarning {
    [super didReceiveMemoryWarning];
    // Dispose of any resources that can be recreated.
}


@end
