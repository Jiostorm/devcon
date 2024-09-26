// Manual Memory Managemnet

const std = @import("std");

const Person = struct {
    name: []const u8,
    age: u8,

    fn init(name: []const u8, age: u8) Person {
        const person = Person{ .name = name, .age = age };
        return person;
    }
};

pub fn main() !void {
    const allocator = std.heap.page_allocator;

    const jio = Person.init("Jio", 23);
    const cla = Person.init("Cla", 23);

    var persons1 = std.ArrayList(Person).init(allocator);
    defer persons1.deinit();
    try persons1.insert(0, jio);
    try persons1.insert(1, cla);
    std.debug.print("{any}\n", .{persons1.items});

    const persons2 = try allocator.alloc(Person, 2);
    defer allocator.free(persons2);
    persons2[0] = jio;
    persons2[1] = cla;
    std.debug.print("{any}\n", .{persons2});

    sayHello(jio);
    sayHello(cla);

    std.debug.print("{s}, {d}\n", .{ jio.name, jio.age });
    std.debug.print("{s}, {d}\n", .{ cla.name, cla.age });
}

fn sayHello(person: Person) void {
    std.debug.print("Hello {s}!\n", .{person.name});
}
