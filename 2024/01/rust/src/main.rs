use std::fs::read_to_string;

fn main() {
    let raw_input = read_to_string("../input.txt").unwrap();
    let mut raw_input: Vec<_> = raw_input
        .trim()
        .split("\n")
        .map(|x| {
            x.split("   ")
                .map(|x| x.trim().parse::<u32>().unwrap())
                .collect::<Vec<_>>()
        })
        .collect();

    let mut lefts: Vec<u32> = vec![];
    let mut rights: Vec<u32> = vec![];

    raw_input.iter_mut().for_each(|line| {
        lefts.push(line[0]);
        rights.push(line[1])
    });

    lefts.sort();
    rights.sort();

    let mut total_distance: u32 = 0;
    for (i, l) in lefts.iter().enumerate() {
        if *l > rights[i] {
            total_distance += l - rights[i];
        } else {
            total_distance += rights[i] - l;
        }
    }

    println!("{:?}", total_distance);
}
