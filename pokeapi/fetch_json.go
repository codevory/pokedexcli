package pokeapi 

import (
	"fmt"
	"net/http"
	"io"
	"context"
	"encoding/json"
	"time"
)

func fetchJSON[T any](context context.Context,c *Client,url string)(T, error,int){
    var res T 

	// cache hit 
	if cacheData,ok := c.cache.Get(url); ok {
		fmt.Printf("[CACHE HIT]  %s (%v)\n", url, time.Since(time.Now()))
		var resp T
        err := json.Unmarshal(cacheData,&resp)
        
		if err != nil {
			return res,err,500
		}

		return resp,nil,200
	}
   

	// cache miss, make an api request
	fmt.Printf("[HTTP GET]   %s\n", url)

	req,err := http.NewRequestWithContext(context,http.MethodGet,url,nil)
	if err != nil {
		return res,fmt.Errorf("build request: %w",err),400
	}

	resp,err := c.httpClient.Do(req)
	if err != nil {
		return res,fmt.Errorf("execute request: %w",err),400
	}

	defer resp.Body.Close()
	fmt.Printf("[HTTP %d]   %s (%v)\n", resp.StatusCode, url, time.Since(time.Now()))

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return res,fmt.Errorf("unexpected status %d from %s",resp.StatusCode,url),resp.StatusCode
	}

	data, err := io.ReadAll(resp.Body)
	 if err != nil {
		return res, err,resp.StatusCode
	 }

	if ok := c.cache.Add(url,data); ok {
		fmt.Sprintf("Data with key %s, added to cache ",url)
	} else {
		fmt.Sprintf("failed to add the data to cache with key %s",url)
	}

	var out T
	if err = json.Unmarshal(data,&out); err != nil {
		return res,fmt.Errorf("decode json %w",err),resp.StatusCode
	}

	return out,nil,resp.StatusCode
}