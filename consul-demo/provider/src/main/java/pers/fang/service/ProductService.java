package pers.fang.service;

import pers.fang.entity.Product;

import java.util.List;

/**
 * @author FANG
 * @create 2022-03-26 16:11
 */
public interface ProductService {
    List<Product> selectProductList();
}
