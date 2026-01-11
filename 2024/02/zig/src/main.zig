const std = @import("std");

pub fn main() !void {
    const file = try std.fs.cwd().openFile("../input.txt", .{});
    const file_reader = file.reader();
    defer file.close();

    var gpa = std.heap.GeneralPurposeAllocator(.{}){};
    const alloc = gpa.allocator();
    defer _ = gpa.deinit();

    var left = std.ArrayList(u32).init(alloc);
    defer left.deinit();

    var right = std.ArrayList(u32).init(alloc);
    defer right.deinit();

    // prob need to use a alloc here too
    var fileBuf: [100_000]u8 = undefined;
    _ = try file_reader.readAll(&fileBuf);

    var lines = std.mem.splitSequence(u8, &fileBuf, "\n");

    while (true) {
        const line = lines.next();

        if (line == null) break;

        var integers = std.mem.splitSequence(u8, line.?, "   ");
        const firstRaw = integers.next();
        const secondRaw = integers.next();

        if (firstRaw == null) {
            continue;
        }

        if (secondRaw == null) {
            continue;
        }

        const first = std.fmt.parseInt(u32, firstRaw.?, 10) catch continue;
        const second = std.fmt.parseInt(u32, secondRaw.?, 10) catch continue;

        try left.append(first);
        try right.append(second);
    }

    std.mem.sort(u32, left.items, {}, std.sort.asc(u32));
    std.mem.sort(u32, right.items, {}, std.sort.asc(u32));

    var totalDistance: usize = 0;
    for (0..left.items.len) |i| {
        totalDistance += if (left.items[i] > right.items[i]) @abs(left.items[i] - right.items[i]) else @abs(right.items[i] - left.items[i]);
    }

    std.debug.print("Total distance {d}\n", .{totalDistance});
}
