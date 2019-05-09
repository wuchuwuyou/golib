//
//  AppDelegate.m
//  test001
//
//  Created by Murphy on 2017/11/28.
//  Copyright © 2017年 Murphy. All rights reserved.
//

#import "AppDelegate.h"
#import "golib.h"
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
    printf("This is a C Application.\n");
    GoString name = {(char*)"King", 4};
    SayHello(name);
    GoSlice buf = {(void*)"King", 4, 4};
    SayHelloByte(buf);
    SayBye();
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
