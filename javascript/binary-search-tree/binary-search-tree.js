class BinarySearchTree {

    constructor(val) {
        this.val = val;
        this.left = this.right = null;
    }

    insert(val) {
        if (val < this.val) {
            if (!this.left) {
                this.left = new BinarySearchTree(val);
            } else {
                this.left.insert(val);
            }
        } else {
            if (!this.right) {
                this.right = new BinarySearchTree(val);
            } else {
                this.right.insert(val);
            }
        }
        return this;
    }

    find(val) {
        if (val === this.val) {
            return this;
        }
        if (val < this.val) {
            if (!this.left) {
                return null;
            }
            return this.left.find(val);
        } else {
            if (!this.right) {
                return null;
            }
            return this.right.find(val);
        }
    }

    traverse(func) {
        if (this.left) {
            this.left.traverse(func);
        }

        func(this.val);

        if (this.right) {
            this.right.traverse(func)
        }
    }
}


module.exports = BinarySearchTree;

/*

const root = new BinarySearchTree(5);

root.insert(7)
    .insert(8)
    .insert(9)
    .insert(10)
    .insert(6);

root.insert(4)
    .insert(2)
    .insert(1)
    .insert(3);

console.log(root);

console.log(root.find(8));
console.log(root.find(12));
console.log(root.traverse());

*/
