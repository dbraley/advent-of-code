const std = @import("std");

pub fn main() !void {
    var gpa = std.heap.GeneralPurposeAllocator(.{}){};
    defer _ = gpa.deinit();
    const allocator = gpa.allocator();

    // Get command line arguments
    const args = try std.process.argsAlloc(allocator);
    defer std.process.argsFree(allocator, args);

    if (args.len < 2) {
        std.debug.print("Please provide a file path\n", .{});
        return;
    }

    // Create ArrayLists to store numbers from each column
    var col1 = std.ArrayList(i32).init(allocator);
    defer col1.deinit();
    var col2 = std.ArrayList(i32).init(allocator);
    defer col2.deinit();

    const file = try std.fs.cwd().openFile(args[1], .{});
    defer file.close();

    var buf_reader = std.io.bufferedReader(file.reader());
    var in_stream = buf_reader.reader();

    // Read the file line by line
    var buf: [1024]u8 = undefined;
    while (try in_stream.readUntilDelimiterOrEof(&buf, '\n')) |line| {
        // Skip empty lines
        if (line.len == 0) continue;
        
        // Split the line and parse numbers, using whitespace as delimiter
        var it = std.mem.tokenize(u8, line, " \t");
        const num1 = try std.fmt.parseInt(i32, it.next() orelse continue, 10);
        const num2 = try std.fmt.parseInt(i32, it.next() orelse continue, 10);

        try col1.append(num1);
        try col2.append(num2);
    }

    // Sort both arrays
    std.sort.heap(i32, col1.items, {}, std.sort.asc(i32));
    std.sort.heap(i32, col2.items, {}, std.sort.asc(i32));

    // Calculate sum of differences
    var total_difference: i32 = 0;
    for (col1.items, col2.items) |v1, v2| {
        total_difference += if (v1 > v2) v1 - v2 else v2 - v1;
    }

    std.debug.print("Total difference: {d}\n", .{total_difference});
}

test "simple test" {
    var list = std.ArrayList(i32).init(std.testing.allocator);
    defer list.deinit(); // try commenting this out and see if zig detects the memory leak!
    try list.append(42);
    try std.testing.expectEqual(@as(i32, 42), list.pop());
}
