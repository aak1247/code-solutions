// Definition for a binary tree node.
// #[derive(Debug, PartialEq, Eq)]
// pub struct TreeNode {
//   pub val: i32,
//   pub left: Option<Rc<RefCell<TreeNode>>>,
//   pub right: Option<Rc<RefCell<TreeNode>>>,
// }
// 
// impl TreeNode {
//   #[inline]
//   pub fn new(val: i32) -> Self {
//     TreeNode {
//       val,
//       left: None,
//       right: None
//     }
//   }
// }
use std::rc::Rc;
use std::cell::RefCell;

impl Solution {
    pub fn max_path_sum(root: Option<Rc<RefCell<TreeNode>>>) -> i32 {
        let mut max_num = i32::min_value();
        let s1 = max_path_sum(&root, &mut max_num);
        max(s1, max_num)
    }
}

pub fn max_path_sum(root: &Option<Rc<RefCell<TreeNode>>>, max_num: &mut i32) -> i32 {
    let val = valOf(root);
    let leftNode = &(root.as_ref()).unwrap().borrow().left;
    let left = if leftNode.is_some() {max_path_sum(leftNode, max_num)} else {i32::min_value()};
    let rightNode = &(root.as_ref()).unwrap().borrow().right;
    let right = if rightNode.is_some() {max_path_sum(rightNode, max_num)} else {i32::min_value()};
    let s1 = max(val+max(0,left), val+max(0,right));
    let s2 = val+ max(0,left)+max(0,right);
    *max_num = max(*max_num, max(s1,s2));
    s1
}

pub fn valOf(node: &Option<Rc<RefCell<TreeNode>>>) -> i32 {
    if node.is_none() { return i32::min_value();}
    ((node.as_ref().unwrap())).borrow().val
}

pub fn max(a:i32,b:i32) -> i32 {
    if a>b {a} else {b}
}