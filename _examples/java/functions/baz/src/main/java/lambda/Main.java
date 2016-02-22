package lambda;

import com.amazonaws.services.lambda.runtime.Context; 

public class Main {

    public static class ExampleRequest {
        String hello;

        public String getHello() {
            return hello;
        }

        public void setHello(String hello) {
            this.hello = hello;
        }

        public ExampleRequest(String hello) {
            this.hello = hello;
        }

        public ExampleRequest() {
        }
    }


    public static class ExampleResponse {
        String hello;

        public String getHello() {
            return hello;
        }

        public void setHello(String hello) {
            this.hello = hello;
        }

        public ExampleResponse(String hello) {
            this.hello = hello;
        }

        public ExampleResponse() {
        }
    }

    public ExampleResponse handler(ExampleRequest event, Context context) {
        return new ExampleResponse("baz");
    }
}
