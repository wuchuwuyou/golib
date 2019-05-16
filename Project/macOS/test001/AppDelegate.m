//
//  AppDelegate.m
//  test001
//
//  Created by Murphy on 2017/11/28.
//  Copyright © 2017年 Murphy. All rights reserved.
//

#import "AppDelegate.h"
#import "golib.h"
#import <Framework/Framework.h>
@interface AppDelegate ()

@property (weak) IBOutlet NSWindow *window;
@end

@implementation AppDelegate

- (void)applicationDidFinishLaunching:(NSNotification *)aNotification {
    // Insert code here to initialize your application
//    NSString *localPath = @"/Users/murphy/Downloads/test/CCC";
//    dispatch_after(dispatch_time(DISPATCH_TIME_NOW, (int64_t)(5* NSEC_PER_SEC)), dispatch_get_main_queue(), ^{
//        NSArray *arr = [self DFSSubPathFromLocalPath:localPath];
//        NSLog(@"%@",arr);
//    });
//
//    NSString *num1 = @"2.0.12345.6";
//    NSString *num2 = @"2.0.2345.6";
//    NSString *num1 = @"2.7.14.2345";
//    NSString *num2 = @"2.12.8.1234";
//    //2.7.14.2345
////    2.12.8.1234
//    if ([num1 compare:num2 options:NSNumericSearch] == NSOrderedDescending) {
//        NSLog(@"%@ is bigger",num1);
//    }else
//    {
//        NSLog(@"%@ is bigger",num2);
//    }
    BOOL isExist = [[NSFileManager defaultManager] fileExistsAtPath:@"/Users/murphy/Desktop/test1.txt"];
    NSLog(@"%@",@(isExist));
    BOOL isEx = NO;
    NSError *error = nil;
    FrameworkIsExist(@"/Users/murphy/Desktop/test1.txt", &isEx, &error);
    NSLog(@"%@,%@",@(isEx),error);
    const char *path1 = [@"/Users/murphy/Desktop/test.txt" UTF8String];
    const char *path2 = [@"/Users/murphy/Desktop/test2.txt" UTF8String];
    const char *path3 = [@"/Users/murphy/Desktop/" UTF8String];
    int s1 = strlen(path1);
    GoString p1 = {path1,30};
    GoString p2 = {path2,strlen(path2)};
    GoString p3 = {path3,strlen(path3)};

//    FileExist(p1);
//    [self print:FileExist(p1)];
//    GoString path = {(char *)"/Users/murphy/Desktop/test.txt",30};
//    IsExist(path);
//    [self print:IsExist(p2).r1];
//
//
//    [self print:ReadDir(p3).r1];
//
//    const char *path4 = [@"/Users/murphy/Desktop/test3.txt" UTF8String];
//    GoString p4 = {path4,strlen(path4)};
//
//    [self print:CreateFile(p4)];

    NSString *username = @"wangteng@lenovocloud.com";
    GoString _username = {[username UTF8String],username.length};
    NSString *password = @"123456";
    GoString _password = {[password UTF8String],password.length};
    dispatch_async(dispatch_get_global_queue(DISPATCH_QUEUE_PRIORITY_DEFAULT, 0), ^{
        GoInterface a = Login(_username, _password);
        [self print:a];
    });
    
}
- (void)print:(GoInterface)a {
    NSLog(@"%@,%@",a.t,a.v);
}
- (NSArray <NSString *>*)DFSSubPathFromLocalPath:(NSString *)localPath {
    NSFileManager *fileManager = [NSFileManager defaultManager];
    BOOL isDir = NO;
    BOOL isExists = [fileManager fileExistsAtPath:localPath isDirectory:&isDir];
    if (isExists == NO) {
        return nil;
    }
    if (isDir == NO) { /// 文件
        return @[localPath];
    }
    NSURL *folderURL = [NSURL URLWithString:localPath];
    NSArray *subFiles = [fileManager contentsOfDirectoryAtURL:folderURL
                                   includingPropertiesForKeys:nil
                                                      options:NSDirectoryEnumerationSkipsHiddenFiles
                                                        error:nil];
    if (subFiles.count == 0) { /// 空文件夹
        return @[[folderURL path]];
    }
    /// 有子文件/子文件夹的情况
    NSMutableArray *array = [NSMutableArray array];
    for (NSURL *subURL in subFiles) {
        NSArray * subArray = [self DFSSubPathFromLocalPath:[subURL path]];
        [array addObjectsFromArray:subArray];
        subArray = nil;
    }
    return array;
}

- (void)applicationWillTerminate:(NSNotification *)aNotification {
    // Insert code here to tear down your application
}


@end
