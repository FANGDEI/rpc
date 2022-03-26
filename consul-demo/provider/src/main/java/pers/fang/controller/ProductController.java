package pers.fang.controller;

import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;
import pers.fang.entity.Product;
import pers.fang.service.ProductService;

import javax.annotation.Resource;
import java.util.List;

/**
 * @author FANG
 * @create 2022-03-26 16:14
 */
@RestController
@RequestMapping("/product")
public class ProductController {
    @Resource
    private ProductService productService;

    @GetMapping("/list")
    public List<Product> selectProductList() {
        return productService.selectProductList();
    }
}
