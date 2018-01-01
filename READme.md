# dot 语法

## 图的属性
1. bgcolor: 设置图的背景颜色，可以使用red、blue、green等形式，或者使用"#rrggbb"形式。

2. label: 设置图的描述。label会继承到子图，子图里如果不想重复显示必须手动设置

3. rankdir: 设置图的方向，包括：TB（top to bottom）、BT（bottom to top）、LR（left to right）、RL（right to left）。

4. rotate: 设置图的旋转。如：rotate = 90便是逆时针旋转90度。

5. ratio: 设置图的长宽比，可以是一个浮点数，也可以是：fill、compress或者auto。

## 结点属性
1. shape: 设置结点形状。包括：Mrecord（圆角矩形）、record（矩形）、circle（圆形）、box（矩形，和record略有区别，下面会讲到）、egg（蛋形）、doublecircle（双圆形）、plaintext（纯文本）、 ellipse（椭圆，默认）。

2. label: 设置结点的显示内容，内容用双引号包含，可以使用转义字符。当结点内容!=结点名时使用

3. style: 设置结点的样式。包括：filled(填充)、dotted（点状边框）、solid（普通边框）、dashed（虚线边框）、bold（边框加粗）、invis（隐形）。

4. color: 设置边框颜色。可以使用单词形式或者"#rrggbb"形式。

5. fillcolor: 设置填充颜色，仅style = filled时有效。

6. width: 设置结点宽度。

7. height: 设置结点高度。

8. peripheries: 设置结点边框个数。

9. fontcolor: 设置结点内容颜色。可以使用单词形式或者"#rrggbb"形式。

## 边属性
1. style: 设置边的形状。包括：solid（实线）、dashed（虚线）、dotted（点线）、bold（加粗）、invis（隐形）。

2. label: 设置边标签。内容用双引号包含，可以使用转义字符。

3. color: 设置边颜色。可以使用单词形式或者"#rrggbb"形式。

4. arrowhead: 设置结点箭头样式。包括：none、empty、odiamond等。

