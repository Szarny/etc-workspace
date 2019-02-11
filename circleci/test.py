import unittest
import calc

class TestCalc(unittest.TestCase):
    def test_calc(self):
        self.assertEqual(calc.add(6, 3), 9)
        self.assertEqual(calc.sub(6, 3), 3)
        self.assertEqual(calc.mul(6, 3), 18)
        self.assertEqual(calc.div(6, 3), 2)

unittest.main()