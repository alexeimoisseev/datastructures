const BinarySearchTree = require('./binary-search-tree');
const root = new BinarySearchTree(5);

test('BinarySearchTree can insert', () => {
    root.insert(3);
    expect(root.left.val).toBe(3);
});


test('BinarySearchTree can find existing', () => {
    const three = root.find(3);
    expect(three.val).toBe(3);

});

test('BinarySearchTree can find null for not existing', () => {
    const nil = root.find(2);
    expect(nil).toBe(null);
})

test('BinarySearchTree can traverse', () => {
    const oneToSix = new BinarySearchTree(1)
        .insert(5)
        .insert(2)
        .insert(4)
        .insert(3)
        .insert(6);

    const arr = [];
    oneToSix.traverse((val) => {
        arr.push(val);
    });

    expect(arr).toEqual([1, 2, 3, 4, 5, 6]);
});
