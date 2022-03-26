package pers.fang.service.impl;

import org.springframework.core.ParameterizedTypeReference;
import org.springframework.http.HttpMethod;
import org.springframework.http.ResponseEntity;
import org.springframework.stereotype.Service;
import org.springframework.web.client.RestTemplate;
import pers.fang.entity.Order;
import pers.fang.entity.Product;
import pers.fang.service.OrderService;

import javax.annotation.Resource;
import java.util.List;

/**
 * @author FANG
 * @create 2022-03-26 16:34
 */
@Service
public class OrderServiceImpl implements OrderService {

    @Resource
    private RestTemplate restTemplate;

    @Override
    public Order selectOrderById(Integer id) {
        return new Order(id, "order-001", "China", 22788D, selectProductListByLoadBalancerAnnotation());
    }

    private List<Product> selectProductListByLoadBalancerAnnotation() {
        ResponseEntity<List<Product>> response = restTemplate.exchange(
                "http://service-provider/product/list",
                HttpMethod.GET,
                null,
                new ParameterizedTypeReference<List<Product>>() {
                }
        );
        return response.getBody();
    }
}
