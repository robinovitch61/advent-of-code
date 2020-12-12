use std::fs;

fn main() {
    // let filename = "../10_test_1.txt";
    // let filename = "../10_test_2.txt";
    let filename = "../10_input.txt";

    let contents = fs::read_to_string(filename).expect("Something went wrong reading the file");
    let mut vals = contents
        .split('\n')
        .filter_map(|x| x.parse::<isize>().ok())
        .collect::<Vec<_>>();

    vals.sort();
    vals.insert(0, 0);
    vals.push(vals.iter().max().expect("No max") + 3);

    let mut num_ones = 0;
    let mut num_threes = 0;
    let zipped = vals.iter().zip(vals.iter().skip(1));
    let diffs = zipped.map(|(x, y)| y - x).collect::<Vec<_>>();

    for diff in diffs.iter() {
        if *diff == 3 {
            num_threes += 1;
        } else if *diff == 1 {
            num_ones += 1
        } else {
            println!("Diff of {} found", diff);
        }
    }

    println!(
        "\nNum ones: {}, num threes: {}, product = {}",
        num_ones,
        num_threes,
        num_ones * num_threes
    );

    let mut path_counts: Vec<i64> = vec![0; (*vals.iter().max().unwrap() + 1) as usize];
    path_counts[0] = 1;
    let valid_diffs = [1, 2, 3];
    for val in vals.iter() {
        for valid_diff in valid_diffs.iter() {
            let previous_jolts = val - valid_diff;
            if previous_jolts >= 0 {
                let num_paths_to_previous_jolts = path_counts[previous_jolts as usize];
                path_counts[*val as usize] += num_paths_to_previous_jolts;
            }
        }
    }

    println!(
        "Num ways to reach {} jolts: {}",
        *vals.iter().max().unwrap() + 1,
        path_counts.last().unwrap()
    )
}
