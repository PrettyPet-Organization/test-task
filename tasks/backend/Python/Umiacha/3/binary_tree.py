from collections import deque


class TreeNode():
    """Бинарное дерево.

    Принцип построения:
    self.left_child.value > self.value >= self.right_child.
    """
    def __init__(self, value,
                 left_child=None, right_child=None):
        self.value = value
        self.left_child = left_child
        self.right_child = right_child

    def make_child(self, new_child):
        """Добавляет новый лист в дерево
        в соответствии с принципом построения.
        """
        if new_child.value >= self.value and not self.right_child:
            self.right_child = new_child
        elif new_child.value >= self.value:
            self.right_child.make_child(new_child)
        if new_child.value < self.value and not self.left_child:
            self.left_child = new_child
        elif new_child.value < self.value:
            self.left_child.make_child(new_child)

    def grow_tree(self, values_arr):
        """Генерирует дерево, по очереди создавая узлы из values_arr
        и соединяя их в соответствии с TreeNode.get_child()."""
        for val in values_arr:
            child = TreeNode(val)
            self.make_child(child)

    def dfs(self, node, goal):
        """Вариант применения метода 'поиска в глубину'."""
        # Базовые случаи. Условия остановки углубления рекурсии.
        if node is None:
            return False
        elif node.value == goal:
            return True

        # Прямой ход рекурсии. Углубление рекурсии.
        result = (node.dfs(node.right_child, goal) 
                  or node.dfs(node.left_child, goal))

        # Обратный ход рекурсии.
        return result

    def bfs(self, goal):
        """Вариант применения метода 'поиска в ширину'."""
        queue = deque()
        queue.append(self)

        # Проходимся по всем узлам.
        while queue:
            curr = queue.popleft()

            # Если нашли подходящий элемент -- выходим.
            if curr.value == goal:
                return True

            if not curr.left_child:
                queue.append(curr.left_child)
            if not curr.right_child:
                queue.append(curr.right_child)

        # Если goal не был найден.
        return False

    def is_value_in_children(self, val, method='bfs'):
        """Осуществляет поиск элемента в графе, начиная с текущего.

        Возвращает True, если элемент найден. Иначе -- False.
        """
        if method == 'dfs':
            return self.dfs(self, val)
        return self.bfs(self, val)

    def pre_order(self, root):
        """Прямой обход бинарного дерева, начиная с root."""
        if root:
            print(root.value, end=' ')
            root.pre_order(root.left_child)
            root.pre_order(root.right_child)

    def in_order(self, root):
        """Центрированный обход бинарного дерева, начиная с root."""
        if root:
            root.pre_order(root.left_child)
            print(root.value, end=' ')
            root.pre_order(root.right_child)

    def post_order(self, root):
        """Обратный обход бинарного дерева, начиная с root."""
        if root:
            root.pre_order(root.left_child)
            root.pre_order(root.right_child)
            print(root.value, end=' ')

    def __str__(self):
        return (
            f'\nTreeNode: {self.value} '
            f'\nLeft child: {self.left_child} '
            f'\nRight child: {self.right_child}.'
        )
