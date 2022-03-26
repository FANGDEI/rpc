package pers.fang.service.impl;

import org.springframework.stereotype.Service;
import pers.fang.entity.Product;
import pers.fang.service.ProductService;

import java.util.Arrays;
import java.util.List;

/**
 * @author FANG
 * @create 2022-03-26 16:11
 */
@Service
public class ProductServiceImpl implements ProductService {
    @Override
    public List<Product> selectProductList() {
        return Arrays.asList(
                new Product(1, "华为电脑", 1, 5800D),
                new Product(2, "联想笔记本", 1, 6888D),
                new Product(3, "小米平板", 5, 2020D)
        );
    }
}
